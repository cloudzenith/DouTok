package followrepo

import (
	"context"
	"errors"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"gorm.io/gen"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) AddFollow(ctx context.Context, userId, targetUserId int64) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		f := &model.Follow{
			ID:           snowflakeutil.GetSnowflakeId(),
			UserID:       userId,
			TargetUserID: targetUserId,
		}
		return tx.WithContext(ctx).Follow.Create(f)
	})
}

func (r *PersistRepository) GetFirstFollowRelation(ctx context.Context, userId, targetUserId int64) (int64, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (int64, error) {
		data, err := tx.WithContext(ctx).Follow.Where(
			query.Q.Follow.UserID.Eq(userId),
			query.Q.Follow.TargetUserID.Eq(targetUserId),
		).First()
		if err != nil {
			return 0, err
		}

		return data.ID, nil
	})
}

func (r *PersistRepository) UpdateRelation2UnDeleted(ctx context.Context, id int64) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := tx.WithContext(ctx).Follow.Where(
			query.Q.Follow.ID.Eq(id),
		).Update(
			query.Q.Follow.IsDeleted, false,
		)
		return err
	})
}

func (r *PersistRepository) RemoveFollow(ctx context.Context, userId, targetUserId int64) error {
	_, err := query.Q.WithContext(ctx).Follow.Where(
		query.Q.Follow.UserID.Eq(userId),
		query.Q.Follow.TargetUserID.Eq(targetUserId),
	).Update(
		query.Q.Follow.IsDeleted, true,
	)
	return err
}

func (r *PersistRepository) parseFollowType(followType int32, userId int64) ([]gen.Condition, error) {
	switch followType {
	case 0:
		return []gen.Condition{
			query.Q.Follow.UserID.Eq(userId),
		}, nil
	case 1:
		return []gen.Condition{
			query.Q.Follow.TargetUserID.Eq(userId),
		}, nil
	case 2:
		return []gen.Condition{
			query.Q.Follow.UserID.Eq(userId),
			query.Q.Follow.TargetUserID.Eq(userId),
		}, nil
	}

	return nil, errors.New("unknown follow type")
}

func (r *PersistRepository) ListFollowing(ctx context.Context, userId int64, followType int32, limit, offset int) ([]int64, error) {
	conditions, err := r.parseFollowType(followType, userId)
	if err != nil {
		return nil, err
	}

	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) ([]int64, error) {
		data, err := tx.WithContext(ctx).Follow.Where(conditions...).Limit(limit).Offset(offset).Find()
		if err != nil {
			return nil, err
		}

		var result []int64
		for _, item := range data {
			switch followType {
			case 0:
				result = append(result, item.TargetUserID)
			case 1:
				result = append(result, item.UserID)
			case 2:
				if item.UserID != userId {
					result = append(result, item.UserID)
				}
				if item.TargetUserID != userId {
					result = append(result, item.TargetUserID)
				}
			}
		}

		return result, nil
	})
}

func (r *PersistRepository) CountFollowing(ctx context.Context, userId int64, followType int32) (int64, error) {
	conditions, err := r.parseFollowType(followType, userId)
	if err != nil {
		return 0, err
	}

	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (int64, error) {
		return tx.Follow.Where(conditions...).Count()
	})
}

func (r *PersistRepository) ListFollowingInGivenList(ctx context.Context, userId int64, targetUserIdList []int64) ([]int64, error) {
	data, err := query.Q.WithContext(ctx).Follow.Select(
		query.Q.Follow.TargetUserID,
	).Where(
		query.Q.Follow.UserID.Eq(userId),
		query.Q.Follow.TargetUserID.In(targetUserIdList...),
		query.Q.Follow.IsDeleted.Is(false),
	).Find()
	if err != nil {
		return nil, err
	}

	var result []int64
	for _, item := range data {
		result = append(result, item.TargetUserID)
	}
	return result, nil
}

var _ repoiface.FollowRepository = (*PersistRepository)(nil)
