package followapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	core *svcoreadapter.Adapter
}

func New(core *svcoreadapter.Adapter) *Application {
	return &Application{
		core: core,
	}
}

func (a *Application) AddFollow(ctx context.Context, request *svapi.AddFollowRequest) (*svapi.AddFollowResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.core.AddFollow(ctx, userId, request.UserId); err != nil {
		log.Context(ctx).Errorf("failed to add follow: %v", err)
		return nil, errorx.New(1, "操作失败")
	}

	return &svapi.AddFollowResponse{}, nil
}

func (a *Application) ListFollowing(ctx context.Context, request *svapi.ListFollowingRequest) (*svapi.ListFollowingResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	resp, err := a.core.ListFollow(ctx, userId, v1.FollowType(request.Type), request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list following: %v", err)
		return nil, errorx.New(1, "获取列表失败")
	}

	userInfoList, err := a.core.GetUserInfoByIdList(ctx, resp.UserIdList)
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to get user info by id list: %v", err)
	}

	userInfoMap := make(map[int64]*v1.User)
	for _, userInfo := range userInfoList {
		userInfoMap[userInfo.Id] = userInfo
	}

	var result []*svapi.FollowUser
	for _, id := range resp.UserIdList {
		userInfo := userInfoMap[id]
		if userInfo == nil {
			continue
		}

		result = append(result, &svapi.FollowUser{
			Id:          userInfo.Id,
			Name:        userInfo.Name,
			Avatar:      userInfo.Avatar,
			IsFollowing: true,
		})
	}

	return &svapi.ListFollowingResponse{
		Users: result,
		Pagination: &svapi.PaginationResponse{
			Page:  resp.Pagination.Page,
			Count: resp.Pagination.Count,
			Total: resp.Pagination.Total,
		},
	}, nil
}

func (a *Application) RemoveFollow(ctx context.Context, request *svapi.RemoveFollowRequest) (*svapi.RemoveFollowResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.core.RemoveFollow(ctx, userId, request.UserId); err != nil {
		log.Context(ctx).Errorf("failed to remove follow: %v", err)
		return nil, errorx.New(1, "操作失败")
	}

	return &svapi.RemoveFollowResponse{}, nil
}

var _ svapi.FollowServiceHTTPServer = (*Application)(nil)
