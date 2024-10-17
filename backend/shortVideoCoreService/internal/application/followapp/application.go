package followapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/followserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	follow followserviceiface.FollowService
	v1.UnimplementedFollowServiceServer
}

func New(follow followserviceiface.FollowService) *Application {
	return &Application{
		follow: follow,
	}
}

func (a *Application) AddFollow(ctx context.Context, request *v1.AddFollowRequest) (*v1.AddFollowResponse, error) {
	if err := a.follow.AddFollow(ctx, request.UserId, request.TargetUserId); err != nil {
		log.Context(ctx).Errorf("failed to add follow: %v", err)
		return &v1.AddFollowResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.AddFollowResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) RemoveFollow(ctx context.Context, request *v1.RemoveFollowRequest) (*v1.RemoveFollowResponse, error) {
	if err := a.follow.RemoveFollow(ctx, request.UserId, request.TargetUserId); err != nil {
		log.Context(ctx).Errorf("failed to remove follow: %v", err)
		return &v1.RemoveFollowResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.RemoveFollowResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) ListFollowing(ctx context.Context, request *v1.ListFollowingRequest) (*v1.ListFollowingResponse, error) {
	data, err := a.follow.ListFollowing(ctx, request.UserId, request.FollowType, request.Pagination)
	if err != nil {
		log.Context(ctx).Errorf("failed to list follow: %v", err)
		return &v1.ListFollowingResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.ListFollowingResponse{
		Meta:       utils.GetSuccessMeta(),
		UserIdList: data.UserIdList,
		Pagination: utils.GetPageResponse(data.Count, request.Pagination.Page, request.Pagination.Size),
	}, nil
}

func (a *Application) IsFollowing(ctx context.Context, request *v1.IsFollowingRequest) (*v1.IsFollowingResponse, error) {
	result, err := a.follow.ListFollowingInGivenList(ctx, request.UserId, request.TargetUserIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to check follow: %v", err)
		return &v1.IsFollowingResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.IsFollowingResponse{
		Meta:          utils.GetSuccessMeta(),
		FollowingList: result,
	}, nil
}

func (a *Application) CountFollow(ctx context.Context, request *v1.CountFollowRequest) (*v1.CountFollowResponse, error) {
	followingNum, followerNum, err := a.follow.CountFollow(ctx, request.UserId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count follow: %v", err)
		return &v1.CountFollowResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CountFollowResponse{
		Meta:           utils.GetSuccessMeta(),
		FollowingCount: followingNum,
		FollowerCount:  followerNum,
	}, nil
}

//var _ v1.FollowServiceServer = (*Application)(nil)
