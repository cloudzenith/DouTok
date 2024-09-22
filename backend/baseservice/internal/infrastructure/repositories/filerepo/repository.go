package filerepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/mysqlx"
	"gorm.io/gorm"
)

// GetTableNameFunc is a function that returns the table name for a file
type GetTableNameFunc func(f *models.File) string

type PersistRepository struct {
	db *gorm.DB
}

func New() *PersistRepository {
	return &PersistRepository{
		db: mysqlx.GetDBClient(context.Background()),
	}
}

func (r *PersistRepository) getTable(file *models.File, method GetTableNameFunc) string {
	return method(file)
}

func (r *PersistRepository) handle(ctx context.Context, tx *gorm.DB, file *models.File, method GetTableNameFunc, op func(f *gorm.DB) error) error {
	db := r.db
	if tx != nil {
		db = tx
	}

	table := r.getTable(file, method)
	return op(db.Table(table).WithContext(ctx))
}

func (r *PersistRepository) GetTx() *gorm.DB {
	return r.db.Begin()
}

func (r *PersistRepository) Add(ctx context.Context, tx *gorm.DB, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, tx, file, method, func(f *gorm.DB) error {
		return f.Create(file).Error
	})
}

func (r *PersistRepository) Load(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, nil, file, method, func(f *gorm.DB) error {
		return f.First(file, "id = ?", file.ID).Error
	})
}

func (r *PersistRepository) LoadUploaded(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, nil, file, method, func(f *gorm.DB) error {
		return f.First(file, "id = ? and uploaded = true", file.ID).Error
	})
}

func (r *PersistRepository) Remove(ctx context.Context, tx *gorm.DB, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, tx, file, method, func(f *gorm.DB) error {
		return f.Delete(file, file.ID).Error
	})
}

func (r *PersistRepository) Update(ctx context.Context, tx *gorm.DB, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, tx, file, method, func(f *gorm.DB) error {
		return f.Updates(file).Error
	})
}

func (r *PersistRepository) LoadByHash(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, nil, file, method, func(f *gorm.DB) error {
		return f.First(file, "hash = ? and uploaded = true", file.Hash).Error
	})
}
