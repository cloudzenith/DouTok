package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) AddFavorite(ctx context.Context, id, userId int64, target v1.FavoriteTarget, favoriteType v1.FavoriteType) error {
	req := &v1.AddFavoriteRequest{
		Id:     id,
		UserId: userId,
		Target: target,
		Type:   favoriteType,
	}

	resp, err := a.favorite.AddFavorite(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveFavorite(ctx context.Context, id, userId int64, target v1.FavoriteTarget, favoriteType v1.FavoriteType) error {
	req := &v1.RemoveFavoriteRequest{
		Id:     id,
		UserId: userId,
		Target: target,
		Type:   favoriteType,
	}

	resp, err := a.favorite.RemoveFavorite(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) IsUserFavoriteVideo(ctx context.Context, userId int64, videoIdList []int64) (map[int64]bool, error) {
	var items []*v1.IsFavoriteRequestItem
	for _, id := range videoIdList {
		items = append(items, &v1.IsFavoriteRequestItem{
			BizId:  id,
			UserId: userId,
		})
	}

	req := &v1.IsFavoriteRequest{
		Items: items,
	}
	resp, err := a.favorite.IsFavorite(ctx, req)
	return respcheck.CheckT[map[int64]bool, *v1.Metadata](
		resp, err,
		func() map[int64]bool {
			result := make(map[int64]bool)
			if len(resp.Result) == 0 {
				return result
			}

			for _, item := range resp.Result {
				result[item.BizId] = item.IsFavorite
			}
			return result
		},
	)
}

func (a *Adapter) ListUserFavoriteVideo(ctx context.Context, userId int64, page, size int32) (*v1.ListFavoriteResponse, error) {
	req := &v1.ListFavoriteRequest{
		FavoriteType:  v1.FavoriteType_FAVORITE,
		AggregateType: v1.FavoriteAggregateType_BY_USER,
		Id:            userId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.favorite.ListFavorite(ctx, req)
	return respcheck.CheckT[*v1.ListFavoriteResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListFavoriteResponse {
			return resp
		},
	)
}

func (a *Adapter) CountFavorite4Video(ctx context.Context, videoIdList []int64) (map[int64]int64, error) {
	req := &v1.CountFavoriteRequest{
		Id:            videoIdList,
		AggregateType: v1.FavoriteAggregateType_BY_VIDEO,
		FavoriteType:  v1.FavoriteType_FAVORITE,
	}

	resp, err := a.favorite.CountFavorite(ctx, req)
	return respcheck.CheckT[map[int64]int64, *v1.Metadata](
		resp, err,
		func() map[int64]int64 {
			result := make(map[int64]int64)
			for _, item := range resp.Items {
				result[item.BizId] = item.Count
			}
			return result
		},
	)
}

func (a *Adapter) CountBeFavoriteNumber4User(ctx context.Context, userId int64) (int64, error) {
	req := &v1.CountFavoriteRequest{
		Id:            []int64{userId},
		AggregateType: v1.FavoriteAggregateType_BY_USER,
		FavoriteType:  v1.FavoriteType_FAVORITE,
	}

	resp, err := a.favorite.CountFavorite(ctx, req)
	return respcheck.CheckT[int64, *v1.Metadata](
		resp, err,
		func() int64 {
			if len(resp.Items) == 0 {
				return 0
			}

			return resp.Items[0].Count
		},
	)
}
