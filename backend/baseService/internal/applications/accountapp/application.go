package accountapp

import (
	"context"
	"errors"

	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/accountserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type AccountApplication struct {
	accountService accountserviceiface.AccountService
}

func New(accountService accountserviceiface.AccountService) *AccountApplication {
	return &AccountApplication{
		accountService: accountService,
	}
}

func (a *AccountApplication) Register(ctx context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	accountId, err := a.accountService.Create(ctx, request.GetMobile(), request.GetEmail(), request.GetPassword())
	if err != nil {
		return &api.RegisterResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.RegisterResponse{
		Meta:      utils.GetSuccessMeta(),
		AccountId: accountId,
	}, nil
}

func (a *AccountApplication) checkPassword(checkFunc func() (int64, error)) (*api.CheckAccountResponse, error) {
	accountId, err := checkFunc()
	if err != nil {
		return &api.CheckAccountResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.CheckAccountResponse{
		Meta:      utils.GetSuccessMeta(),
		AccountId: accountId,
	}, nil
}

func (a *AccountApplication) CheckAccount(ctx context.Context, request *api.CheckAccountRequest) (*api.CheckAccountResponse, error) {
	if request.GetAccountId() != 0 {
		return a.checkPassword(func() (int64, error) {
			return a.accountService.CheckPasswordById(ctx, request.GetAccountId(), request.GetPassword())
		})
	}

	if request.GetMobile() != "" {
		return a.checkPassword(func() (int64, error) {
			return a.accountService.CheckPasswordByMobile(ctx, request.GetMobile(), request.GetPassword())
		})
	}

	if request.GetEmail() != "" {
		return a.checkPassword(func() (int64, error) {
			return a.accountService.CheckPasswordByEmail(ctx, request.GetEmail(), request.GetPassword())
		})
	}

	log.Context(ctx).Error("unknown request type")
	return &api.CheckAccountResponse{
		Meta: utils.GetMetaWithError(errors.New("unknown request type")),
	}, nil
}

func (a *AccountApplication) Bind(ctx context.Context, request *api.BindRequest) (*api.BindResponse, error) {
	if err := a.accountService.Bind(ctx, request.GetAccountId(), request.GetVoucherType(), request.GetVoucher()); err != nil {
		log.Context(ctx).Errorf("bind voucher failed: %v", err)
		return &api.BindResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.BindResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *AccountApplication) Unbind(ctx context.Context, request *api.UnbindRequest) (*api.UnbindResponse, error) {
	if err := a.accountService.Unbind(ctx, request.GetAccountId(), request.GetVoucherType()); err != nil {
		log.Context(ctx).Errorf("unbind voucher failed: %v", err)
		return &api.UnbindResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.UnbindResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.AccountServiceServer = (*AccountApplication)(nil)
