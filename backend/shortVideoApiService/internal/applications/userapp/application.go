package userapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter/accountoptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/useroptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/transport"
	jwtv5 "github.com/golang-jwt/jwt/v5"
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

func (a *Application) GetUserInfo(ctx context.Context, request *svapi.GetUserInfoRequest) (*svapi.GetUserInfoResponse, error) {
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

	return &svapi.GetUserInfoResponse{
		User: &svapi.User{
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

func (a *Application) GetVerificationCode(ctx context.Context, request *svapi.GetVerificationCodeRequest) (*svapi.GetVerificationCodeResponse, error) {
	codeId, err := a.base.CreateVerificationCode(ctx, 6, 60*10)
	if err != nil {
		log.Context(ctx).Error("failed to create verification code")
		return nil, errorx.New(1, "failed to get verification code")
	}

	return &svapi.GetVerificationCodeResponse{
		CodeId: codeId,
	}, nil
}

func (a *Application) setToken2Header(ctx context.Context, claim *claims.Claims) error {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	if header, ok := transport.FromServerContext(ctx); ok {
		header.ReplyHeader().Set("Authorization", "Bearer "+tokenString)
	}

	return jwt.ErrWrongContext
}

func (a *Application) Login(ctx context.Context, request *svapi.LoginRequest) (*svapi.LoginResponse, error) {
	userId, err := a.base.CheckAccount(ctx)
	if err != nil {
		log.Context(ctx).Error("failed to check account: %v", err)
		return nil, errorx.New(1, "failed to check account")
	}

	a.setToken2Header(ctx, claims.New(userId))
	return &svapi.LoginResponse{}, nil
}

func (a *Application) Register(ctx context.Context, request *svapi.RegisterRequest) (*svapi.RegisterResponse, error) {
	if err := a.base.ValidateVerificationCode(ctx, request.CodeId, request.Code); err != nil {
		return nil, errorx.New(1, "invalid verification code")
	}

	var options []accountoptions.RegisterOptions
	if request.Mobile != "" {
		options = append(options, accountoptions.RegisterWithMobile(request.Mobile))
	}

	if request.Email != "" {
		options = append(options, accountoptions.RegisterWithEmail(request.Email))
	}

	if request.Password != "" {
		options = append(options, accountoptions.RegisterWithPassword(request.Password))
	}

	userId, err := a.base.Register(ctx, options...)
	if err != nil {
		log.Context(ctx).Error("failed to register account")
		return nil, errorx.New(1, "failed to register account")
	}

	// TODO: 调用core服务创建基本用户信息

	return &svapi.RegisterResponse{
		UserId: userId,
	}, nil
}

func (a *Application) UpdateUserInfo(ctx context.Context, request *svapi.UpdateUserInfoRequest) (*svapi.UpdateUserInfoResponse, error) {
	if err := a.core.UpdateUserInfo(
		ctx,
		useroptions.UpdateUserInfoWithUserId(request.UserId),
		useroptions.UpdateUserInfoWithName(request.Name),
		useroptions.UpdateUserInfoWithAvatar(request.Avatar),
		useroptions.UpdateUserInfoWithBackgroundImage(request.BackgroundImage),
		useroptions.UpdateUserInfoWithSignature(request.Signature),
	); err != nil {
		log.Context(ctx).Error("failed to update user info")
		return nil, errorx.New(1, "failed to update user info")
	}

	return &svapi.UpdateUserInfoResponse{}, nil
}

var _ svapi.UserServiceHTTPServer = (*Application)(nil)
