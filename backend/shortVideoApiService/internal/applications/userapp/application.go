package userapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New("unknown user info")
	}

	return &api.GetUserInfoResponse{
		User: &api.User{
			Name: "test",
		},
	}, errorx.New(21, "test error")
}

func (a *Application) GetVerificationCode(ctx context.Context, request *api.GetVerificationCodeRequest) (*api.GetVerificationCodeResponse, error) {
	return nil, nil
}

func (a *Application) Login(ctx context.Context, request *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}

func (a *Application) Register(ctx context.Context, request *api.RegisterRequest) (*api.RegisterResponse, error) {
	return nil, nil
}

func (a *Application) UpdateUserInfo(ctx context.Context, request *api.UpdateUserInfoRequest) (*api.UpdateUserInfoResponse, error) {
	return nil, nil
}

var _ api.UserServiceHTTPServer = (*Application)(nil)
