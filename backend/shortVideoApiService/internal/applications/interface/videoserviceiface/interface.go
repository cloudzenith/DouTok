package videoserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

type VideoService interface {
	AssembleVideo(ctx context.Context, userId int64, data []*v1.Video) ([]*svapi.Video, error)
	AssembleAuthorInfo(ctx context.Context, data []*svapi.Video)
	AssembleVideoList(ctx context.Context, userId int64, data []*v1.Video) ([]*svapi.Video, error)
	AssembleUserIsFollowing(ctx context.Context, list []*svapi.Video, userId int64)
	AssembleVideoCountInfo(ctx context.Context, list []*svapi.Video)
}
