package userservice

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	config      *conf.Config
	log         *log.Helper
	userUsecase UserUsecase
	v1.UnimplementedUserServiceServer
}

func NewUserService(config *conf.Config, logger log.Logger, user UserUsecase) *UserService {
	return &UserService{
		config:      config,
		log:         log.NewHelper(logger),
		userUsecase: user,
	}
}

func (s *UserService) Register(ctx context.Context, in *v1.RegisterRequest) (*v1.RegisterResponse, error) {
	userId, err := s.userUsecase.Register(ctx, in.Mobile, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	return &v1.RegisterResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		UserId: userId,
	}, nil
}

func (s *UserService) Login(ctx context.Context, in *v1.LoginRequest) (*v1.LoginResponse, error) {
	token, err := s.userUsecase.Login(ctx, in.Mobile, in.Email, in.Password)
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

func (s *UserService) UpdateUserInfo(ctx context.Context, in *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoResponse, error) {
	err := s.userUsecase.UpdateUserInfo(ctx, &entity.User{
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

func (s *UserService) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	user, err := s.userUsecase.GetUserInfo(ctx, in.UserId)
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoResponse{
		User: user.ToUserResp(),
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
	}, nil
}
