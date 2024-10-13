package repoiface

import (
	"context"
)

type FollowRepository interface {
	AddFollow(ctx context.Context, userId, targetUserId int64) error
	GetFirstFollowRelation(ctx context.Context, userId, targetUserId int64) (int64, error)
	UpdateRelation2UnDeleted(ctx context.Context, id int64) error
	RemoveFollow(ctx context.Context, userId, targetUserId int64) error
	ListFollowing(ctx context.Context, userId int64, followType int32, limit, offset int) ([]int64, error)
	CountFollowing(ctx context.Context, userId int64, followType int32) (int64, error)
	ListFollowingInGivenList(ctx context.Context, userId int64, targetUserIdList []int64) ([]int64, error)
}
