package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type FavoriteRepository interface {
	AddFavorite(ctx context.Context, userId, targetId int64, targetType, favoriteType int32) error
	RemoveFavorite(ctx context.Context, userId, targetId int64, targetType, favoriteType int32) error
	ListFavorite(ctx context.Context, bizId int64, aggType, favoriteType int32, limit, offset int) ([]int64, error)
	CountFavorite(ctx context.Context, bizId []int64, aggType, favoriteType int32) ([]*CountFavoriteResult, error)
	Get4IsFavorite(ctx context.Context, userId, bizId []int64) ([]*model.Favorite, error)
}

type CountFavoriteResult struct {
	Id  int64
	Cnt int64
}
