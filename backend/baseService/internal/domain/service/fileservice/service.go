package fileservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/file"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/slicingfile"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/innerservice/filerepohelper"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minio/minio-go/v7"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

type FileService struct {
	fileRepo       repoiface.FileRepository
	minio          repoiface.MinioRepository
	config         filerepohelper.FileTableShardingConfig
	fileRepoHelper *filerepohelper.FileRepoHelper
}

func New(fileRepo repoiface.FileRepository, minio repoiface.MinioRepository, config filerepohelper.FileTableShardingConfig) *FileService {
	return &FileService{
		fileRepo:       fileRepo,
		minio:          minio,
		config:         config,
		fileRepoHelper: filerepohelper.New(fileRepo, config),
	}
}

func (s *FileService) CheckFileExistedAndGetFile(ctx context.Context, fileCtx *api.FileContext) (int64, bool, error) {
	f := file.NewWithFileContext(fileCtx)
	fm := f.ToModel()
	err := s.fileRepo.LoadByHash(ctx, fm, s.fileRepoHelper.GetTableNameByHash())
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, false, err
	}

	if err != nil {
		return -1, false, nil
	}

	return fm.ID, true, nil
}

func (s *FileService) PreSignGet(ctx context.Context, fileCtx *api.FileContext) (string, error) {
	f := file.NewWithFileContext(fileCtx)
	fm := f.ToModel()
	if err := s.fileRepo.LoadUploaded(ctx, fm, s.fileRepoHelper.GetTableNameById()); err != nil {
		return "", err
	}

	f = file.NewWithModel(fm)
	return s.minio.PreSignGetUrl(ctx, f.DomainName, f.GetObjectName(), fileCtx.Filename, fileCtx.ExpireSeconds)
}

func (s *FileService) PreSignPut(ctx context.Context, fileCtx *api.FileContext) (string, int64, error) {
	f := file.NewWithFileContext(fileCtx)
	f.SetId()
	fm := f.ToModel()
	if err := s.fileRepoHelper.Add(ctx, fm); err != nil {
		return "", 0, err
	}

	url, err := s.minio.PreSignPutUrl(ctx, f.DomainName, f.GetObjectName(), fileCtx.ExpireSeconds)
	if err != nil {
		return "", 0, err
	}

	return url, f.ID, nil
}

func (s *FileService) PreSignSlicingPut(ctx context.Context, fileCtx *api.FileContext) (*slicingfile.SlicingFile, error) {
	id, ok, err := s.CheckFileExistedAndGetFile(ctx, fileCtx)
	if err != nil {
		log.Context(ctx).Warnf("failed to check file existed, err: %v", err)
	}

	if ok {
		return slicingfile.New(file.NewFile(file.WithID(id))), nil
	}

	f := file.NewFile(
		file.WithDomain(fileCtx.Domain),
		file.WithBizName(fileCtx.BizName),
		file.WithHash(fileCtx.Hash),
		file.WithFileType(fileCtx.FileType),
		file.WithSize(fileCtx.Size),
	)
	f.SetId()
	fm := f.ToModel()
	if err := s.fileRepoHelper.Add(ctx, fm); err != nil {
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
	if err := s.fileRepo.Load(ctx, fm, s.fileRepoHelper.GetTableNameById()); err != nil {
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

func (s *FileService) ReportUploaded(ctx context.Context, fileCtx *api.FileContext) error {
	fm := &models.File{}
	fm.ID = fileCtx.FileId
	fm.DomainName = fileCtx.Domain
	fm.BizName = fileCtx.BizName
	if err := s.fileRepo.Load(ctx, fm, s.fileRepoHelper.GetTableNameById()); err != nil {
		return err
	}

	f := file.NewWithModel(fm)
	hash, err := s.minio.GetObjectHash(ctx, fm.DomainName, f.GetObjectName())
	if err != nil {
		return err
	}

	equals := f.CheckHash(hash)
	if !equals && !strings.Contains(hash, "-") {
		log.Context(ctx).Errorf("failed to validate hash of uploaded file, hash: %s, expected: %s", hash, f.Hash)
		return errors.New("failed to validate hash of uploaded file")
	}

	return s.fileRepoHelper.Update(ctx, f.SetUploaded().ToModel())
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
	if err := s.fileRepo.Load(ctx, fm, s.fileRepoHelper.GetTableNameById()); err != nil {
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
	if err := s.fileRepoHelper.Remove(ctx, fm); err != nil {
		return err
	}

	return nil
}

func (s *FileService) GetInfoById(ctx context.Context, domainName, bizName string, fileId int64) (*file.File, error) {
	fm := &models.File{}
	fm.ID = fileId
	fm.DomainName = domainName
	fm.BizName = bizName
	if err := s.fileRepo.Load(ctx, fm, s.fileRepoHelper.GetTableNameById()); err != nil {
		return nil, err
	}

	return file.NewWithModel(fm), nil
}
