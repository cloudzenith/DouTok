package accountapp

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/accountserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
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
	if err := validate(request); err != nil {
		return &api.RegisterResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

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

func (a *AccountApplication) checkPassword(ctx context.Context, checkFunc func() (int64, error)) (*api.CheckAccountResponse, error) {
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
	if err := validate(request); err != nil {
		return &api.CheckAccountResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	if request.GetAccountId() != 0 {
		return a.checkPassword(ctx, func() (int64, error) {
			return a.accountService.CheckPasswordById(ctx, request.GetAccountId(), request.GetPassword())
		})
	}

	if request.GetMobile() != "" {
		return a.checkPassword(ctx, func() (int64, error) {
			return a.accountService.CheckPasswordByMobile(ctx, request.GetMobile(), request.GetPassword())
		})
	}

	if request.GetEmail() != "" {
		return a.checkPassword(ctx, func() (int64, error) {
			return a.accountService.CheckPasswordByEmail(ctx, request.GetEmail(), request.GetPassword())
		})
	}

	return &api.CheckAccountResponse{
		Meta: utils.GetMetaWithError(errors.New("unknown request type")),
	}, nil
}

var _ api.AccountServiceServer = (*AccountApplication)(nil)
