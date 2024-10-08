package favoriteapp

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

var _ svapi.FavoriteServiceHTTPServer = (*Application)(nil)
