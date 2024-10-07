package userapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/userdomain"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type UserApplication struct {
	userUsecase userdomain.IUserUsecase
	v1.UnimplementedUserServiceServer
}

func NewUserApplication(user userdomain.IUserUsecase) *UserApplication {
	return &UserApplication{
		userUsecase: user,
	}
}

func (s *UserApplication) CreateUser(ctx context.Context, in *v1.CreateUserRequest) (*v1.CreateUserResponse, error) {
	userId, err := s.userUsecase.CreateUser(ctx, in.Mobile, in.Email, in.AccountId)
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		UserId: userId,
	}, nil
}

func (s *UserApplication) UpdateUserInfo(ctx context.Context, in *v1.UpdateUserInfoRequest) (*v1.UpdateUserInfoResponse, error) {
	log.Context(ctx).Infof("UpdateUserInfo: %v", in)
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
			BizCode: 0,
			Message: "success",
		},
	}, nil
}

func (s *UserApplication) GetUserInfo(ctx context.Context, in *v1.GetUserInfoRequest) (*v1.GetUserInfoResponse, error) {
	user, err := s.userUsecase.GetUserInfo(ctx, dto.GetUserInfoRequest{
		UserId:    in.UserId,
		AccountId: in.AccountId,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GetUserInfoResponse{
		User: user.ToUserResp(),
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
	}, nil
}

func (s *UserApplication) GetUserByIdList(ctx context.Context, in *v1.GetUserByIdListRequest) (*v1.GetUserByIdListResponse, error) {
	data, err := s.userUsecase.GetUserByIdList(ctx, in.UserIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user by id list: %v", err)
		return &v1.GetUserByIdListResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var users []*v1.User
	for _, user := range data {
		users = append(users, user.ToUserResp())
	}

	return &v1.GetUserByIdListResponse{
		Meta:     utils.GetSuccessMeta(),
		UserList: users,
	}, nil
}
