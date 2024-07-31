package userdomain

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
)

type UserRepo interface {
	Save(ctx context.Context, u *model.User) error
	UpdateById(ctx context.Context, u *model.User) (int64, error)
	FindByID(ctx context.Context, id int64) (*model.User, error)
	FindByAccountID(ctx context.Context, accountID int64) (*model.User, error)
}

var _ UserRepo = (*data.UserRepo)(nil)
