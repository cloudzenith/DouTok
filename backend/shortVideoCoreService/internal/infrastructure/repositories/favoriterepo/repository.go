package favoriterepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) AddFavorite(ctx context.Context, userId, targetId int64, targetType, favoriteType int32) error {
	f := &model.Favorite{
		ID:           snowflakeutil.GetSnowflakeId(),
		UserID:       userId,
		TargetID:     targetId,
		TargetType:   targetType,
		FavoriteType: favoriteType,
		IsDeleted:    false,
	}
	return query.Q.WithContext(ctx).Favorite.Create(f)
}

func (r *PersistRepository) RemoveFavorite(ctx context.Context, userId, targetId int64, targetType, favoriteType int32) error {
	_, err := query.Q.WithContext(ctx).Favorite.Where(
		query.Q.Favorite.UserID.Eq(userId),
		query.Q.Favorite.TargetID.Eq(targetId),
		query.Q.Favorite.TargetType.Eq(targetType),
		query.Q.Favorite.FavoriteType.Eq(favoriteType),
	).Update(query.Q.Favorite.IsDeleted, true)
	return err
}

func (r *PersistRepository) ListFavorite(ctx context.Context, bizId int64, aggType, favoriteType int32, limit, offset int) ([]int64, error) {
	var conditions []gen.Condition
	if aggType == int32(v1.FavoriteAggregateType_BY_USER) {
		conditions = append(conditions, query.Q.Favorite.UserID.Eq(bizId))
	} else {
		conditions = append(conditions, query.Q.Favorite.TargetID.Eq(bizId))
	}

	conditions = append(conditions, query.Q.Favorite.FavoriteType.Eq(favoriteType))
	result, err := query.Q.WithContext(ctx).Favorite.Where(
		conditions...,
	).Limit(limit).Offset(offset).Find()
	if err != nil {
		return nil, err
	}

	var res []int64
	for _, item := range result {
		if aggType == int32(v1.FavoriteAggregateType_BY_USER) {
			res = append(res, item.TargetID)
		} else {
			res = append(res, item.TargetID)
		}
	}

	return res, nil
}

func (r *PersistRepository) CountFavorite(ctx context.Context, bizId []int64, aggType, favoriteType int32) ([]*repoiface.CountFavoriteResult, error) {
	var fields []field.Expr
	var targetField field.Expr
	var result []*repoiface.CountFavoriteResult
	if aggType == int32(v1.FavoriteAggregateType_BY_USER) {
		targetField = query.Q.Favorite.UserID
		fields = append(fields, query.Q.Favorite.UserID.As("id"))
		fields = append(fields, query.Q.Favorite.UserID.Count().As("cnt"))
	} else {
		targetField = query.Q.Favorite.TargetID
		fields = append(fields, query.Q.Favorite.TargetID.As("id"))
		fields = append(fields, query.Q.Favorite.TargetID.Count().As("cnt"))
	}

	err := query.Q.WithContext(ctx).Favorite.Select(
		fields...,
	).Where(
		query.Q.Favorite.TargetID.In(bizId...),
	).Group(
		targetField,
	).Scan(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *PersistRepository) Get4IsFavorite(ctx context.Context, userId, bizId []int64) ([]*model.Favorite, error) {
	return query.Q.WithContext(ctx).Favorite.Where(
		query.Q.Favorite.UserID.In(userId...),
		query.Q.Favorite.TargetID.In(bizId...),
		query.Q.Favorite.IsDeleted.Is(true),
	).Find()
}

var _ repoiface.FavoriteRepository = (*PersistRepository)(nil)
