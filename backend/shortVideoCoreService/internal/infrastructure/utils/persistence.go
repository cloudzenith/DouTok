package utils

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"gorm.io/gorm"
)

type DBClient struct {
	db *gorm.DB
}

func (c *DBClient) GetDB() *gorm.DB {
	return c.db
}

type Transaction interface {
	ExecTx(context.Context, func(ctx context.Context) error) error
}

type contextTxKey struct{}

// ExecTx gorm Transaction
func (c *DBClient) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTxKey{}, tx)
		return fn(ctx)
	})
}

func (c *DBClient) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return c.db
}

type WhereConditionFn func(db *gorm.DB) *gorm.DB

func (c *DBClient) WhereWithPaginateAndSort(
	ctx context.Context,
	fn WhereConditionFn,
	value interface{},
	sort string,
	request *dto.PaginationRequest,
) *gorm.DB {
	offset := (request.PageNum - 1) * request.PageSize
	db := fn(c.DB(ctx)).Offset(int(offset)).Limit(int(request.PageSize)).Order(sort).Find(value)
	return db
}
