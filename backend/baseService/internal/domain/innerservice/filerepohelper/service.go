package filerepohelper

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/filerepo"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type FileTableShardingConfig interface {
	GetShardingNumber(tableName, domainName, bizName string) int64
}

type FileRepoHelper struct {
	fileRepo repoiface.FileRepository
	config   FileTableShardingConfig
}

func New(fileRepo repoiface.FileRepository, config FileTableShardingConfig) *FileRepoHelper {
	return &FileRepoHelper{
		fileRepo: fileRepo,
		config:   config,
	}
}

func (h *FileRepoHelper) GetTableNameById() filerepo.GetTableNameFunc {
	return func(f *models.File) string {
		shardingNum := h.config.GetShardingNumber(models.TableNameFile, f.DomainName, f.BizName)

		if shardingNum > 0 {
			return fmt.Sprintf("%s_%s_%s_id_%d", models.TableNameFile, f.DomainName, f.BizName, f.ID%shardingNum)
		}

		return fmt.Sprintf("%s_%s", f.DomainName, f.BizName)
	}
}

func (h *FileRepoHelper) GetTableNameByHash() filerepo.GetTableNameFunc {
	return func(f *models.File) string {
		shardingNum := h.config.GetShardingNumber(models.TableNameFile, f.DomainName, f.BizName)

		if shardingNum > 0 {
			return fmt.Sprintf("%s_%s_%s_hash_%d", models.TableNameFile, f.DomainName, f.BizName, int64([]byte(f.Hash)[0])%shardingNum)
		}

		return fmt.Sprintf("%s_%s", f.DomainName, f.BizName)
	}
}

func (h *FileRepoHelper) doWithTx(
	ctx context.Context, file *models.File, opName string,
	method func(c context.Context, t *gorm.DB, f *models.File, getTableFunc filerepo.GetTableNameFunc) error,
) (err error) {
	if file.ID == 0 || len(file.Hash) == 0 {
		return errors.New("file can't be do something with tx if without id and hash")
	}

	tx := h.fileRepo.GetTx()
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Context(ctx).Errorf("failed to roallback saving file: %v", file)
			}

			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			log.Context(ctx).Errorf("failed to commit saving file: %v", file)
		}
	}()

	if tx == nil {
		log.Context(ctx).Errorf("failed to begin tx when save file: %v", file)
		return errors.New("failed to begin tx when save file")
	}

	err = method(ctx, tx, file, h.GetTableNameById())
	if err != nil {
		log.Context(ctx).Errorf("failed to %s with id: %v", opName, file)
		return err
	}

	err = method(ctx, tx, file, h.GetTableNameByHash())
	if err != nil {
		log.Context(ctx).Errorf("failed to %s with hash: %v", opName, file)
		return err
	}

	return nil
}

func (h *FileRepoHelper) Add(ctx context.Context, file *models.File) (err error) {
	return h.doWithTx(ctx, file, "save file", h.fileRepo.Add)
}

func (h *FileRepoHelper) Remove(ctx context.Context, file *models.File) error {
	return h.doWithTx(ctx, file, "remove file", h.fileRepo.Remove)
}

func (h *FileRepoHelper) Update(ctx context.Context, file *models.File) error {
	return h.doWithTx(ctx, file, "update file", h.fileRepo.Update)
}
