package favoriteapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/interface/videoserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	core         *svcoreadapter.Adapter
	videoService videoserviceiface.VideoService
}

func New(core *svcoreadapter.Adapter, videoService videoserviceiface.VideoService) *Application {
	return &Application{
		core:         core,
		videoService: videoService,
	}
}

func (a *Application) AddFavorite(ctx context.Context, request *svapi.AddFavoriteRequest) (*svapi.AddFavoriteResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.core.AddFavorite(ctx, request.Id, userId, v1.FavoriteTarget(request.Target), v1.FavoriteType(request.Type)); err != nil {
		log.Context(ctx).Errorf("failed to add favorite: %v", err)
		return nil, errorx.New(1, "操作失败")
	}

	return &svapi.AddFavoriteResponse{}, nil
}

func (a *Application) RemoveFavorite(ctx context.Context, request *svapi.RemoveFavoriteRequest) (*svapi.RemoveFavoriteResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.core.RemoveFavorite(ctx, request.Id, userId, v1.FavoriteTarget(request.Target), v1.FavoriteType(request.Type)); err != nil {
		log.Context(ctx).Errorf("failed to remove favorite: %v", err)
		return nil, errorx.New(1, "操作失败")
	}

	return &svapi.RemoveFavoriteResponse{}, nil
}

func (a *Application) ListFavoriteVideo(ctx context.Context, request *svapi.ListFavoriteVideoRequest) (*svapi.ListFavoriteVideoResponse, error) {
	if request.UserId == 0 {
		userId, err := claims.GetUserId(ctx)
		if err != nil {
			return nil, errorx.New(1, "获取用户信息失败")
		}
		request.UserId = userId
	}

	resp, err := a.core.ListUserFavoriteVideo(ctx, request.UserId, request.Page, request.Size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list favorite video: %v", err)
		return nil, errorx.New(1, "获取喜欢列表失败")
	}

	if len(resp.BizId) == 0 {
		return &svapi.ListFavoriteVideoResponse{
			Videos: make([]*svapi.Video, 0),
			Pagination: &svapi.PaginationResponse{
				Page:  request.Page,
				Total: 0,
				Count: 0,
			},
		}, nil
	}

	videoList, err := a.core.GetVideosByIdList(ctx, resp.BizId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get videos by id list: %v", err)
		return nil, errorx.New(1, "获取喜欢列表失败")
	}

	result, err := a.videoService.AssembleVideo(ctx, request.UserId, videoList)
	if err != nil {
		log.Context(ctx).Warnf("something wrong in assembling videos: %v", err)
	}

	return &svapi.ListFavoriteVideoResponse{
		Videos:     result,
		Pagination: respcheck.ParseSvCorePagination(resp.Pagination),
	}, nil
}

var _ svapi.FavoriteServiceHTTPServer = (*Application)(nil)
