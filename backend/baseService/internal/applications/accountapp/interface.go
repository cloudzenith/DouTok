package accountapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/accountservice"
)

type AccountService interface {
	Create(ctx context.Context, mobile, email, password string) (int64, error)
	CheckPasswordById(ctx context.Context, id int64, password string) (int64, error)
	CheckPasswordByMobile(ctx context.Context, mobile, password string) (int64, error)
	CheckPasswordByEmail(ctx context.Context, email, password string) (int64, error)
	ModifyPassword(ctx context.Context, id int64, oldPassword, newPassword string) error
}

var _ AccountService = (*accountservice.Service)(nil)
