package followservice

import (
	"context"
	"errors"
	"github.com/TremblingV5/box/dbtx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/followserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type Service struct {
	follow repoiface.FollowRepository
}

func New(follow repoiface.FollowRepository) *Service {
	return &Service{
		follow: follow,
	}
}

func (s *Service) AddFollow(ctx context.Context, userId, targetUserId int64) (err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	existedRelationId, err := s.follow.GetFirstFollowRelation(ctx, userId, targetUserId)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Context(ctx).Errorf("failed to get first follow relation: %v", err)
		return err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.follow.AddFollow(ctx, userId, targetUserId)
		if err != nil {
			log.Context(ctx).Errorf("failed to add follow relation: %v", err)
			return err
		}

		return err
	}

	err = s.follow.UpdateRelation2UnDeleted(ctx, existedRelationId)
	if err != nil {
		log.Context(ctx).Errorf("failed to update relation to un-deleted: %v", err)
		return err
	}

	return nil
}

func (s *Service) RemoveFollow(ctx context.Context, userId, targetUserId int64) error {
	if err := s.follow.RemoveFollow(ctx, userId, targetUserId); err != nil {
		log.Context(ctx).Errorf("failed to remove follow relation: %v", err)
		return err
	}

	return nil
}

func (s *Service) ListFollowing(ctx context.Context, userId int64, followType v1.FollowType, pagination *v1.PaginationRequest) (result *followserviceiface.ListFollowingDTO, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	var limit = int(pagination.Size)
	var offset = (int(pagination.Page) - 1) * int(pagination.Size)
	data, err := s.follow.ListFollowing(ctx, userId, int32(followType), limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("failed to list following: %v", err)
		return nil, err
	}

	count, err := s.follow.CountFollowing(ctx, userId, int32(followType))
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to count following: %v", err)
	}

	return &followserviceiface.ListFollowingDTO{
		UserIdList: data,
		Count:      count,
	}, nil
}

func (s *Service) ListFollowingInGivenList(ctx context.Context, userId int64, targetUserIdList []int64) ([]int64, error) {
	result, err := s.follow.ListFollowingInGivenList(ctx, userId, targetUserIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to list following in given list: %v", err)
		return nil, err
	}

	return result, nil
}

func (s *Service) CountFollow(ctx context.Context, userId int64) (followingNum int64, followerNum int64, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	followingNum, err = s.follow.CountFollowing(ctx, userId, int32(v1.FollowType_FOLLOWING))
	if err != nil {
		log.Context(ctx).Errorf("failed to count following: %v", err)
		return 0, 0, err
	}

	followerNum, err = s.follow.CountFollowing(ctx, userId, int32(v1.FollowType_FOLLOWER))
	if err != nil {
		log.Context(ctx).Errorf("failed to count follower: %v", err)
		return 0, 0, err
	}

	return followingNum, followerNum, nil
}

var _ followserviceiface.FollowService = (*Service)(nil)
