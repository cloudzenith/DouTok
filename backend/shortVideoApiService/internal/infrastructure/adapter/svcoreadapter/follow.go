package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

func (a *Adapter) AddFollow(ctx context.Context, userId, targetUserId int64) error {
	req := &v1.AddFollowRequest{
		UserId:       userId,
		TargetUserId: targetUserId,
	}

	resp, err := a.follow.AddFollow(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveFollow(ctx context.Context, userId, targetUserId int64) error {
	req := &v1.RemoveFollowRequest{
		UserId:       userId,
		TargetUserId: targetUserId,
	}

	resp, err := a.follow.RemoveFollow(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) ListFollow(ctx context.Context, userId int64, followType v1.FollowType, page, size int32) (*v1.ListFollowingResponse, error) {
	req := &v1.ListFollowingRequest{
		UserId:     userId,
		FollowType: followType,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.follow.ListFollowing(ctx, req)
	return respcheck.CheckT[*v1.ListFollowingResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListFollowingResponse {
			return resp
		},
	)
}

func (a *Adapter) IsFollowing(ctx context.Context, userId int64, targetUserIdList []int64) (map[int64]bool, error) {
	req := &v1.IsFollowingRequest{
		UserId:           userId,
		TargetUserIdList: targetUserIdList,
	}

	resp, err := a.follow.IsFollowing(ctx, req)
	log.Context(ctx).Infof("IsFollowing resp: %v", resp)
	return respcheck.CheckT[map[int64]bool, *v1.Metadata](
		resp, err,
		func() map[int64]bool {
			result := make(map[int64]bool)
			if len(resp.FollowingList) == 0 {
				return result
			}

			for _, item := range resp.FollowingList {
				result[item] = true
			}
			return result
		},
	)
}

func (a *Adapter) CountFollow4User(ctx context.Context, userId int64) ([]int64, error) {
	req := &v1.CountFollowRequest{
		UserId: userId,
	}

	resp, err := a.follow.CountFollow(ctx, req)
	return respcheck.CheckT[[]int64, *v1.Metadata](
		resp, err,
		func() []int64 {
			return []int64{resp.FollowingCount, resp.FollowerCount}
		},
	)
}
