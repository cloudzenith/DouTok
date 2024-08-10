package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/filerepo"
)

//go:generate mockgen -source=file.go -destination=file_mock.go -package=repoiface FileRepository
type FileRepository interface {
	Add(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	Load(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	Remove(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
}
