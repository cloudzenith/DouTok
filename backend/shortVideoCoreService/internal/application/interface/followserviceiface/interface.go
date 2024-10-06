package followserviceiface

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

type ListFollowingDTO struct {
	Count      int64
	UserIdList []int64
}

type FollowService interface {
	AddFollow(ctx context.Context, userId, targetUserId int64) error
	RemoveFollow(ctx context.Context, userId, targetUserId int64) error
	ListFollowing(ctx context.Context, userId int64, followType v1.FollowType, pagination *v1.PaginationRequest) (*ListFollowingDTO, error)
}
