package userdomain

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
)

type IUserUsecase interface {
	CreateUser(ctx context.Context, mobile, email string, accountId int64) (int64, error)
	GetUserInfo(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUserInfo(ctx context.Context, user *entity.User) error
}

var _ IUserUsecase = (*UserUsecase)(nil)
