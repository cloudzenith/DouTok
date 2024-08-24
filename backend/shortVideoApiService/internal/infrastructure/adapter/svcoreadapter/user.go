package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/useroptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) GetUserInfo(ctx context.Context, userId int64) (*v1.User, error) {
	req := &v1.GetUserInfoRequest{
		UserId: userId,
	}

	resp, err := a.user.GetUserInfo(ctx, req)
	return respcheck.CheckT[*v1.User, *v1.Metadata](
		resp, err,
		func() *v1.User {
			return resp.User
		},
	)
}

func (a *Adapter) UpdateUserInfo(ctx context.Context, options ...useroptions.UpdateUserInfoOption) error {
	req := &v1.UpdateUserInfoRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.user.UpdateUserInfo(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}
