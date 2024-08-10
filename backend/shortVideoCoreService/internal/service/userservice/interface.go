package userservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/userdomain"
)

type UserUsecase interface {
	Register(ctx context.Context, mobile, email, password string) (int64, error)
	Login(ctx context.Context, mobile, email, password string) (string, error)
	GetUserInfo(ctx context.Context, userId int64) (*entity.User, error)
	UpdateUserInfo(ctx context.Context, user *entity.User) error
}

var _ UserUsecase = (*userdomain.UserUsecase)(nil)
