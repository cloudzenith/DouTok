package userapp

import (
	"context"
	"fmt"
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

func (a *Application) checkUserId(ctx context.Context, receivedUserId int64) (int64, error) {
	if receivedUserId != 0 {
		return receivedUserId, nil
	}

	return claims.GetUserId(ctx)
}

func (a *Application) GetUserInfo(ctx context.Context, request *svapi.GetUserInfoRequest) (resp *svapi.GetUserInfoResponse, err error) {
	userId, err := a.checkUserId(ctx, request.UserId)
	if err != nil {
		return nil, errorx.New(1, "failed to parse user id when getting user info from token")
	}

	userInfo, err := a.core.GetUserInfo(ctx, useroptions.GetUserInfoWithUserId(userId))
	if err != nil {
		log.Context(ctx).Error("failed to get user info")
		log.Context(ctx).Errorw("error", err, "user_id", request.UserId)
		return nil, errorx.New(1, "failed to get user info")
	}

	result := &svapi.User{
		Id:              userInfo.Id,
		Name:            userInfo.Name,
		Avatar:          userInfo.Avatar,
		BackgroundImage: userInfo.BackgroundImage,
		Signature:       userInfo.Signature,
		Mobile:          userInfo.Mobile,
		Email:           userInfo.Email,
		WorkCount:       userInfo.WorkCount,
		FavoriteCount:   userInfo.FavoriteCount,
	}

	followCount, err := a.core.CountFollow4User(ctx, userId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count follow: %v", err)
	}
	if err == nil && len(followCount) == 2 {
		result.FollowCount = followCount[0]
		result.FollowerCount = followCount[1]
	}

	beFavoriteCount, err := a.core.CountBeFavoriteNumber4User(ctx, userId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count be favorite: %v", err)
	}
	if err == nil {
		result.TotalFavorited = beFavoriteCount
	}

	return &svapi.GetUserInfoResponse{
		User: result,
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

func (a *Application) setToken2Header(ctx context.Context, claim *claims.Claims) (string, error) {
	token := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claim)
	tokenString, err := token.SignedString([]byte("token"))
	if err != nil {
		return "", err
	}

	if header, ok := transport.FromServerContext(ctx); ok {
		header.ReplyHeader().Set("Authorization", "Bearer "+tokenString)
		return tokenString, nil
	}

	return "", jwt.ErrWrongContext
}

func (a *Application) Login(ctx context.Context, request *svapi.LoginRequest) (*svapi.LoginResponse, error) {
	accountId, err := a.base.CheckAccount(
		ctx,
		accountoptions.CheckAccountWithMobile(request.GetMobile()),
		accountoptions.CheckAccountWithEmail(request.GetEmail()),
		accountoptions.CheckAccountWithPassword(request.GetPassword()),
	)
	if err != nil {
		log.Context(ctx).Error("failed to check account: %v", err)
		return nil, errorx.New(1, "failed to check account")
	}
	user, err := a.core.GetUserInfo(ctx, useroptions.GetUserInfoWithAccountId(accountId))
	if err != nil {
		log.Context(ctx).Error("failed to get user info: %v", err)
		return nil, errorx.New(1, "failed to get user info")
	}

	token, err := a.setToken2Header(ctx, claims.New(user.Id))
	if err != nil {
		log.Context(ctx).Error("failed to generate token: %v", err)
		return nil, errorx.New(1, "failed to generate token")
	}
	return &svapi.LoginResponse{
		Token: token,
	}, nil
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

	accountId, err := a.base.Register(ctx, options...)
	if err != nil {
		log.Context(ctx).Error("failed to register account")
		return nil, errorx.New(1, "failed to register account")
	}

	// TODO: 调用core服务创建基本用户信息, 需要处理 register 成功，但是创建用户信息失败
	userId, err := a.core.CreateUser(ctx, request.Mobile, request.Email, accountId)
	if err != nil {
		log.Context(ctx).Error(fmt.Sprintf("failed to create user: %v", err))
		return nil, errorx.New(1, fmt.Sprintf("failed to create user: %v", err))
	}
	return &svapi.RegisterResponse{
		UserId: userId,
	}, nil
}

func (a *Application) UpdateUserInfo(ctx context.Context, request *svapi.UpdateUserInfoRequest) (*svapi.UpdateUserInfoResponse, error) {
	log.Context(ctx).Infof("UpdateUserInfo: %v", request)
	userId, err := a.checkUserId(ctx, request.UserId)
	if err != nil {
		return nil, errorx.New(1, "failed to parse user id when getting user info from token")
	}

	if err := a.core.UpdateUserInfo(
		ctx,
		useroptions.UpdateUserInfoWithUserId(userId),
		useroptions.UpdateUserInfoWithName(request.Name),
		useroptions.UpdateUserInfoWithAvatar(request.Avatar),
		useroptions.UpdateUserInfoWithBackgroundImage(request.BackgroundImage),
		useroptions.UpdateUserInfoWithSignature(request.Signature),
	); err != nil {
		log.Context(ctx).Errorf("failed to update user info: %v", err)
		return nil, errorx.New(1, "failed to update user info")
	}

	return &svapi.UpdateUserInfoResponse{}, nil
}

func (a *Application) BindUserVoucher(ctx context.Context, request *svapi.BindUserVoucherRequest) (*svapi.BindUserVoucherResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) UnbindUserVoucher(ctx context.Context, request *svapi.UnbindUserVoucherRequest) (*svapi.UnbindUserVoucherResponse, error) {
	//TODO implement me
	panic("implement me")
}

var _ svapi.UserServiceHTTPServer = (*Application)(nil)
