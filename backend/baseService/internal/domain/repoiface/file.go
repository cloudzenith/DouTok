package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/filerepo"
	"gorm.io/gorm"
)

//go:generate mockgen -source=file.go -destination=file_mock.go -package=repoiface FileRepository
type FileRepository interface {
	GetTx() *gorm.DB
	Add(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
	Load(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	LoadByHash(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	Remove(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
	Update(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
}
