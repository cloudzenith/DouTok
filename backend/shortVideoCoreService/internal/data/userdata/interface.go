package userdata

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
)

type IUserRepo interface {
	Save(ctx context.Context, tx *query.Query, u *model.User) error
	UpdateById(ctx context.Context, tx *query.Query, u *model.User) (int64, error)
	FindByID(ctx context.Context, tx *query.Query, id int64) (*model.User, error)
	FindByAccountID(ctx context.Context, tx *query.Query, accountID int64) (*model.User, error)
	FindByIds(ctx context.Context, tx *query.Query, ids []int64) ([]*model.User, error)
}

var _ IUserRepo = (*UserRepo)(nil)
