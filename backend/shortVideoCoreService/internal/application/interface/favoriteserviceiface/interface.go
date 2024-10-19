package favoriteserviceiface

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils/pageresult"
)

type FavoriteService interface {
	AddFavorite(ctx context.Context, dto *WriteOpDTO) error
	RemoveFavorite(ctx context.Context, dto *WriteOpDTO) error
	ListFavorite(ctx context.Context, dto *AggOpDTO, limit, offset int) (*pageresult.R[int64], error)
	CountFavorite(ctx context.Context, dto *AggOpDTO) ([]*v1.CountFavoriteResponseItem, error)
	IsFavorite(ctx context.Context, dto []*v1.IsFavoriteRequestItem) ([]*v1.IsFavoriteResponseItem, error)
}

type WriteOpDTO struct {
	UserId       int64
	TargetId     int64
	TargetType   v1.FavoriteTarget
	FavoriteType v1.FavoriteType
}

func (dto *WriteOpDTO) Check() error {
	//if &dto.TargetType == nil || &dto.FavoriteType == nil {
	//	return errors.New("查询类型错误")
	//}

	return nil
}

type AggOpDTO struct {
	BizId        int64
	BizIdList    []int64
	AggType      v1.FavoriteAggregateType
	FavoriteType v1.FavoriteType
}

func (dto *AggOpDTO) Check() error {
	//if &dto.AggType == nil || &dto.FavoriteType == nil {
	//	return errors.New("查询类型错误")
	//}

	return nil
}
