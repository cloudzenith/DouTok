package service

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/do"
	"github.com/go-kratos/kratos/v2/log"
)

type ShortVideoCoreService struct {
	log  *log.Helper
	user *domain.UserUsecase
	v1.UnimplementedShortVideoCoreUserServiceServer
}

func NewShortVideoCoreService(
	user *domain.UserUsecase,
	logger log.Logger) *ShortVideoCoreService {
	return &ShortVideoCoreService{user: user, log: log.NewHelper(logger)}
}

func (s *ShortVideoCoreService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	user, err := s.user.Register(ctx, &do.User{
		Mobile:   in.Mobile,
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.RegisterResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		UserId: user.ID,
	}, nil
}

func (s *ShortVideoCoreService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	token, err := s.user.Login(ctx, &do.User{
		Mobile:   in.Mobile,
		Email:    in.Email,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}
	return &v1.LoginResponse{
		Token: token,
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}

func (s *ShortVideoCoreService) UpdateUserInfo(ctx context.Context, in *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoResponse, error) {
	_, err := s.user.UpdateUserInfo(ctx, &do.User{
		ID:              in.UserId,
		Name:            in.Name,
		Avatar:          in.Avatar,
		BackgroundImage: in.BackgroundImage,
		Signature:       in.Signature,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserInfoResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}

func (s *ShortVideoCoreService) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	user, err := s.user.GetUserInfo(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoResponse{
		User: &v1.User{
			Id:              user.ID,
			Name:            user.Name,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			Mobile:          user.Mobile,
			Email:           user.Email,
		},
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}
