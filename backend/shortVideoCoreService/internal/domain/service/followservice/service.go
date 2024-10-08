package followservice

import (
	"context"
	"github.com/TremblingV5/box/dbtx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/followserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	follow repoiface.FollowRepository
}

func New(follow repoiface.FollowRepository) *Service {
	return &Service{
		follow: follow,
	}
}

func (s *Service) AddFollow(ctx context.Context, userId, targetUserId int64) error {
	if err := s.follow.AddFollow(ctx, userId, targetUserId); err != nil {
		log.Context(ctx).Errorf("failed to add follow relation: %v", err)
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

var _ followserviceiface.FollowService = (*Service)(nil)
