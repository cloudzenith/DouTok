package favoriteapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/favoriteserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	service favoriteserviceiface.FavoriteService
	v1.UnimplementedFavoriteServiceServer
}

func New(favoriteService favoriteserviceiface.FavoriteService) *Application {
	return &Application{
		service: favoriteService,
	}
}

func (a *Application) AddFavorite(ctx context.Context, request *v1.AddFavoriteRequest) (*v1.AddFavoriteResponse, error) {
	if err := a.service.AddFavorite(ctx, &favoriteserviceiface.WriteOpDTO{
		UserId:       request.UserId,
		TargetId:     request.Id,
		FavoriteType: request.Type,
		TargetType:   request.Target,
	}); err != nil {
		return &v1.AddFavoriteResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.AddFavoriteResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) RemoveFavorite(ctx context.Context, request *v1.RemoveFavoriteRequest) (*v1.RemoveFavoriteResponse, error) {
	if err := a.service.RemoveFavorite(ctx, &favoriteserviceiface.WriteOpDTO{
		UserId:       request.UserId,
		TargetId:     request.Id,
		FavoriteType: request.Type,
		TargetType:   request.Target,
	}); err != nil {
		return &v1.RemoveFavoriteResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.RemoveFavoriteResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) ListFavorite(ctx context.Context, request *v1.ListFavoriteRequest) (*v1.ListFavoriteResponse, error) {
	data, err := a.service.ListFavorite(ctx, &favoriteserviceiface.AggOpDTO{
		BizId:        request.Id,
		AggType:      request.AggregateType,
		FavoriteType: request.FavoriteType,
	},
		int(request.Pagination.Size),
		(int(request.Pagination.Page)-1)*int(request.Pagination.Size),
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to list favorites: %v", err)
		return &v1.ListFavoriteResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.ListFavoriteResponse{
		Meta:       utils.GetSuccessMeta(),
		BizId:      data.Data,
		Pagination: utils.GetPageResponse(data.Count, request.Pagination.Page, request.Pagination.Size),
	}, nil
}

func (a *Application) CountFavorite(ctx context.Context, request *v1.CountFavoriteRequest) (*v1.CountFavoriteResponse, error) {
	data, err := a.service.CountFavorite(ctx, &favoriteserviceiface.AggOpDTO{
		BizIdList:    request.Id,
		AggType:      request.AggregateType,
		FavoriteType: request.FavoriteType,
	})
	if err != nil {
		log.Context(ctx).Errorf("failed to count favorites: %v", err)
		return &v1.CountFavoriteResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CountFavoriteResponse{
		Meta:  utils.GetSuccessMeta(),
		Items: data,
	}, nil
}

func (a *Application) IsFavorite(ctx context.Context, request *v1.IsFavoriteRequest) (*v1.IsFavoriteResponse, error) {
	data, err := a.service.IsFavorite(ctx, request.Items)
	if err != nil {
		log.Context(ctx).Errorf("failed to check is favorite: %v", err)
		return &v1.IsFavoriteResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.IsFavoriteResponse{
		Meta:   utils.GetSuccessMeta(),
		Result: data,
	}, nil
}

//var _ v1.FavoriteServiceServer = (*Application)(nil)
