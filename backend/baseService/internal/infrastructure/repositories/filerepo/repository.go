package filerepo

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/golang/groupcache/singleflight"
	"gorm.io/gorm"
)

type GetTableNameFunc func(f *models.File) string

type PersistRepository struct {
	db                      *gorm.DB
	createTableSingleFlight singleflight.Group
}

func New(db *gorm.DB) *PersistRepository {
	return &PersistRepository{
		db:                      db,
		createTableSingleFlight: singleflight.Group{},
	}
}

func (r *PersistRepository) getTable(file *models.File, method GetTableNameFunc) (string, error) {
	tableName := method(file)

	_, err := r.createTableSingleFlight.Do(tableName, func() (interface{}, error) {
		isExists := r.db.Migrator().HasTable(tableName)
		if isExists {
			return tableName, nil
		}

		return nil, r.db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `%s` LIKE `%s`", tableName, models.TableNameFile)).Error
	})
	if err != nil {
		return "", nil
	}

	return tableName, nil
}

func (r *PersistRepository) handle(ctx context.Context, file *models.File, method GetTableNameFunc, op func(f *gorm.DB) error) error {
	table, err := r.getTable(file, method)
	if err != nil {
		return err
	}

	return op(r.db.Table(table).WithContext(ctx))
}

func (r *PersistRepository) Add(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, file, method, func(f *gorm.DB) error {
		return f.Create(file).Error
	})
}

func (r *PersistRepository) Load(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, file, method, func(f *gorm.DB) error {
		return f.First(file, file.ID).Error
	})
}

func (r *PersistRepository) Remove(ctx context.Context, file *models.File, method GetTableNameFunc) error {
	return r.handle(ctx, file, method, func(f *gorm.DB) error {
		return f.Delete(file, file.ID).Error
	})
}
