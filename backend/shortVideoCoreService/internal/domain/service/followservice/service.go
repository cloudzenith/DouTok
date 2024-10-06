package followservice

import (
	"context"
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

func (s *Service) ListFollowing(ctx context.Context, userId int64, followType v1.FollowType, pagination *v1.PaginationRequest) ([]int64, error) {
	var limit = int(pagination.Size)
	var offset = (int(pagination.Page) - 1) * int(pagination.Size)
	data, err := s.follow.ListFollowing(ctx, userId, int32(followType), limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("failed to list following: %v", err)
		return nil, err
	}

	return data, nil
}

var _ followserviceiface.FollowService = (*Service)(nil)
