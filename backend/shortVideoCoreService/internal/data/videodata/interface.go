package videodata

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
)

type IVideoRepo interface {
	Save(ctx context.Context, v *model.Video) error
	UpdateById(ctx context.Context, v *model.Video) (int64, error)
	FindByID(ctx context.Context, id int64) (*model.Video, error)
	GetVideoList(ctx context.Context, request *dto.GetVideoListRequest) (*dto.GetVideoListResponse, error)
	GetVideoFeed(ctx context.Context, request *dto.GetVideoFeedRequest) (*dto.GetVideoFeedResponse, error)
}

var _ IVideoRepo = (*VideoRepo)(nil)
