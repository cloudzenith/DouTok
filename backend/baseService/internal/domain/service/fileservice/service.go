package fileservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/file"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/slicingfile"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/filerepo"
	"github.com/minio/minio-go/v7"
	"strconv"
)

type FileTableShardingConfig interface {
	GetShardingNumber(domainName, bizName string) int64
}

type FileService struct {
	fileRepo repoiface.FileRepository
	minio    repoiface.MinioRepository
	config   FileTableShardingConfig
}

func New(fileRepo repoiface.FileRepository, minio repoiface.MinioRepository, config FileTableShardingConfig) *FileService {
	return &FileService{
		fileRepo: fileRepo,
		minio:    minio,
		config:   config,
	}
}

func (s *FileService) getTableName() filerepo.GetTableNameFunc {
	return func(f *models.File) string {
		shardingNum := s.config.GetShardingNumber(f.DomainName, f.BizName)

		if shardingNum > 0 {
			return fmt.Sprintf("%s_%s_%d", f.DomainName, f.BizName, f.ID%shardingNum)
		}

		return fmt.Sprintf("%s_%s", f.DomainName, f.BizName)
	}
}

func (s *FileService) PreSignGet(ctx context.Context, fileCtx *api.FileContext) (string, error) {
	f := file.NewWithFileContext(fileCtx)
	fm := f.ToModel()
	if err := s.fileRepo.Load(ctx, fm, s.getTableName()); err != nil {
		return "", err
	}

	f = file.NewWithModel(fm)
	return s.minio.PreSignGetUrl(ctx, f.DomainName, f.GetObjectName(), fileCtx.Filename, fileCtx.ExpireSeconds)
}

func (s *FileService) PreSignPut(ctx context.Context, fileCtx *api.FileContext) (string, error) {
	f := file.NewWithFileContext(fileCtx)
	f.SetId()
	fm := f.ToModel()
	if err := s.fileRepo.Add(ctx, fm, s.getTableName()); err != nil {
		return "", err
	}

	return s.minio.PreSignPutUrl(ctx, f.DomainName, f.GetObjectName(), fileCtx.ExpireSeconds)
}

func (s *FileService) PreSignSlicingPut(ctx context.Context, fileCtx *api.FileContext) (*slicingfile.SlicingFile, error) {
	f := file.NewFile(
		file.WithDomain(fileCtx.Domain),
		file.WithBizName(fileCtx.BizName),
		file.WithHash(fileCtx.Hash),
		file.WithFileType(fileCtx.FileType),
		file.WithSize(fileCtx.Size),
	)
	f.SetId()
	fm := f.ToModel()
	if err := s.fileRepo.Add(ctx, fm, s.getTableName()); err != nil {
		return nil, err
	}

	uploadId, err := s.minio.CreateSlicingUpload(ctx, f.DomainName, f.GetObjectName(), minio.PutObjectOptions{})
	if err != nil {
		return nil, err
	}

	slicingFile := slicingfile.New(f).SetUploadId(uploadId).SetTotalParts()
	urls := make([]string, slicingFile.TotalParts)
	for i := 1; i <= int(slicingFile.TotalParts); i++ {
		url, e := s.minio.PreSignSlicingPutUrl(ctx, f.DomainName, f.GetObjectName(), uploadId, int64(i))
		if e != nil {
			return nil, e
		}
		urls[i-1] = url
	}

	slicingFile.UploadUrl = urls
	return slicingFile, nil
}

func (s *FileService) checkSlicingFileUploaded(res map[string]bool) (bool, string) {
	total := 0
	finished := 0
	for _, uploaded := range res {
		if uploaded {
			finished++
		}

		total++
	}

	rate := fmt.Sprintf("%d/%d", finished, total)
	return total == finished, rate

}

func (s *FileService) GetProgressRate4SlicingPut(ctx context.Context, uploadId string, fileCtx *api.FileContext) (map[string]bool, error) {
	fm := &models.File{}
	fm.ID = fileCtx.FileId
	fm.DomainName = fileCtx.Domain
	fm.BizName = fileCtx.BizName
	if err := s.fileRepo.Load(ctx, fm, s.getTableName()); err != nil {
		return nil, err
	}

	f := file.NewWithModel(fm)
	sf := slicingfile.New(f).SetTotalParts()
	result, err := s.minio.ListSlicingFileParts(ctx, f.DomainName, f.GetObjectName(), uploadId, sf.TotalParts)
	if err != nil {
		return nil, err
	}

	res := make(map[string]bool)
	parts := result.ObjectParts
	for i := 0; i < int(sf.TotalParts); i++ {
		if len(parts[i].ETag) > 0 {
			res[strconv.FormatInt(int64(i+1), 10)] = true
		} else {
			res[strconv.FormatInt(int64(i+1), 10)] = false
		}
	}

	return res, nil
}

func (s *FileService) ReportUploadedFileParts(ctx context.Context, uploadId string, fileId, partNumber int64) error {
	return nil
}

func (s *FileService) MergeFileParts(ctx context.Context, uploadId string, fileCtx *api.FileContext) error {
	uploadResult, err := s.GetProgressRate4SlicingPut(ctx, uploadId, fileCtx)
	if err != nil {
		return err
	}

	if ok, _ := s.checkSlicingFileUploaded(uploadResult); !ok {
		return errors.New("not all parts uploaded")
	}

	fm := &models.File{}
	fm.ID = fileCtx.FileId
	fm.DomainName = fileCtx.Domain
	fm.BizName = fileCtx.BizName
	if err := s.fileRepo.Load(ctx, fm, s.getTableName()); err != nil {
		return err
	}

	f := file.NewWithModel(fm)
	sf := slicingfile.New(f).SetTotalParts()

	result, err := s.minio.ListSlicingFileParts(ctx, f.DomainName, f.GetObjectName(), uploadId, sf.TotalParts)
	if err != nil {
		return err
	}

	parts := make([]minio.CompletePart, 0)
	for i := 0; i < len(result.ObjectParts); i++ {
		parts = append(parts, minio.CompletePart{
			PartNumber: i + 1,
			ETag:       result.ObjectParts[i].ETag,
		})
	}

	return s.minio.MergeSlices(ctx, f.DomainName, f.GetObjectName(), uploadId, parts)
}

func (s *FileService) RemoveFile(ctx context.Context, fileCtx *api.FileContext) error {
	f := file.NewWithFileContext(fileCtx)
	fm := f.ToModel()
	if err := s.fileRepo.Remove(ctx, fm, s.getTableName()); err != nil {
		return err
	}

	return nil
}
