package userapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) GetUserInfo(context.Context, *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	return &api.GetUserInfoResponse{
		User: &api.User{
			Name: "test",
		},
	}, errorx.New(21, "test error")
}

func (a *Application) GetVerificationCode(context.Context, *api.GetVerificationCodeRequest) (*api.GetVerificationCodeResponse, error) {
	return nil, nil
}

func (a *Application) Login(context.Context, *api.LoginRequest) (*api.LoginResponse, error) {
	return nil, nil
}

func (a *Application) Register(context.Context, *api.RegisterRequest) (*api.RegisterResponse, error) {
	return nil, nil
}

func (a *Application) UpdateUserInfo(context.Context, *api.UpdateUserInfoRequest) (*api.UpdateUserInfoResponse, error) {
	return nil, nil
}

var _ api.UserServiceHTTPServer = (*Application)(nil)
