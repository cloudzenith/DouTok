package accountserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/accountservice"
)

//go:generate mockgen -source=interface.go -destination=interface_mock.go -package=accountserviceiface AccountService
type AccountService interface {
	Create(ctx context.Context, mobile, email, password string) (int64, error)
	CheckPasswordById(ctx context.Context, id int64, password string) (int64, error)
	CheckPasswordByMobile(ctx context.Context, mobile, password string) (int64, error)
	CheckPasswordByEmail(ctx context.Context, email, password string) (int64, error)
	ModifyPassword(ctx context.Context, id int64, oldPassword, newPassword string) error
	Unbind(ctx context.Context, id int64, voucherType api.VoucherType) error
	Bind(ctx context.Context, id int64, voucherType api.VoucherType, voucher string) error
}

var _ AccountService = (*accountservice.Service)(nil)
