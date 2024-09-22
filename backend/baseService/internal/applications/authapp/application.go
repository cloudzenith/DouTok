package authapp

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/authserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/verificationcode"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type AuthApplication struct {
	authService authserviceiface.AuthService
}

func New(authService authserviceiface.AuthService) *AuthApplication {
	return &AuthApplication{
		authService: authService,
	}
}

func (a *AuthApplication) CreateVerificationCode(ctx context.Context, request *api.CreateVerificationCodeRequest) (*api.CreateVerificationCodeResponse, error) {
	code, err := a.authService.CreateVerificationCode(ctx, request.Bits, request.ExpireTime)
	if err != nil {
		log.Context(ctx).Errorf("failed to create verification code: %v", err)
		return &api.CreateVerificationCodeResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.CreateVerificationCodeResponse{
		Meta:               utils.GetSuccessMeta(),
		VerificationCodeId: code.VerificationCodeId,
	}, nil
}

func (a *AuthApplication) ValidateVerificationCode(ctx context.Context, request *api.ValidateVerificationCodeRequest) (*api.ValidateVerificationCodeResponse, error) {
	code := verificationcode.New(request.VerificationCodeId, request.Code)
	if err := code.IsReady(); err != nil {
		log.Context(ctx).Errorf("verification code info is incomplete: %v", err)
		return &api.ValidateVerificationCodeResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	ok, err := a.authService.ValidateVerificationCode(ctx, code)
	if err != nil {
		log.Context(ctx).Errorf("failed to validate verification code: %v", err)
		return &api.ValidateVerificationCodeResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil

	}

	if !ok {
		return &api.ValidateVerificationCodeResponse{
			Meta: utils.GetMetaWithError(errors.New("verification code is invalid")),
		}, nil
	}

	return &api.ValidateVerificationCodeResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.AuthServiceServer = (*AuthApplication)(nil)
