package userapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	base *baseadapter.Adapter
	core *svcoreadapter.Adapter
}

func New(
	base *baseadapter.Adapter,
	core *svcoreadapter.Adapter,
) *Application {
	return &Application{
		base: base,
		core: core,
	}
}

func (a *Application) GetUserInfo(ctx context.Context, request *api.GetUserInfoRequest) (*api.GetUserInfoResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "unknown user info")
	}

	userInfo, err := a.core.GetUserInfo(ctx, userId)
	if err != nil {
		log.Context(ctx).Error("failed to get user info")
		log.Context(ctx).Errorw("error", err, "user_id", userId)
		return nil, errorx.New(1, "failed to get user info")
	}

	return &api.GetUserInfoResponse{
		User: &api.User{
			Id:              userInfo.Id,
			Name:            userInfo.Name,
			Avatar:          userInfo.Avatar,
			BackgroundImage: userInfo.BackgroundImage,
			Signature:       userInfo.Signature,
			Mobile:          userInfo.Mobile,
			Email:           userInfo.Email,
			FollowCount:     userInfo.FollowCount,
			FollowerCount:   userInfo.FollowerCount,
			TotalFavorited:  userInfo.TotalFavorited,
			WorkCount:       userInfo.WorkCount,
			FavoriteCount:   userInfo.FavoriteCount,
		},
	}, nil
}

func (a *Application) GetVerificationCode(ctx context.Context, request *api.GetVerificationCodeRequest) (*api.GetVerificationCodeResponse, error) {
	codeId, err := a.base.CreateVerificationCode(ctx, 6, 60*10)
	if err != nil {
		log.Context(ctx).Error("failed to create verification code")
		return nil, errorx.New(1, "failed to get verification code")
	}
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
