package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/filerepo"
	"gorm.io/gorm"
)

//go:generate mockgen -source=file.go -destination=file_mock.go -package=repoiface FileRepository
type FileRepository interface {
	// GetTx returns a transaction for file tables
	GetTx() *gorm.DB
	// Add adds a file to the database
	Add(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
	// Load loads a file from the database
	Load(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	// LoadUploaded loads an uploaded file from the database
	LoadUploaded(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	// LoadByHash loads a file by hash from the database
	LoadByHash(ctx context.Context, file *models.File, method filerepo.GetTableNameFunc) error
	// Remove removes a file from the database
	Remove(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
	// Update updates a file in the database
	Update(ctx context.Context, tx *gorm.DB, file *models.File, method filerepo.GetTableNameFunc) error
}
