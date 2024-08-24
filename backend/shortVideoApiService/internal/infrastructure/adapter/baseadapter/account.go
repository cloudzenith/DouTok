package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter/accountoptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
)

func (a *Adapter) Register(ctx context.Context, options ...accountoptions.RegisterOptions) (int64, error) {
	req := &api.RegisterRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.account.Register(ctx, req)
	return respcheck.CheckT[int64, *api.Metadata](
		resp, err,
		func() int64 {
			return resp.AccountId
		},
	)
}

func (a *Adapter) CheckAccount(ctx context.Context, options ...accountoptions.CheckAccountOption) (int64, error) {
	req := &api.CheckAccountRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.account.CheckAccount(ctx, req)
	return respcheck.CheckT[int64, *api.Metadata](
		resp, err,
		func() int64 {
			return resp.AccountId
		},
	)
}
