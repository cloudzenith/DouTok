package videodomain

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
)

type VideoUsecase interface {
	FeedShortVideo(ctx context.Context, request *dto.FeedShortVideoRequest) (*dto.FeedShortVideoResponse, error)
	GetVideoById(ctx context.Context, videoId int64) (*entity.Video, error)
	GetVideoByIdList(ctx context.Context, videoIdList []int64) ([]*entity.Video, error)
	PublishVideo(ctx context.Context, video *dto.PublishVideoRequest) (int64, error)
	ListPublishedVideo(ctx context.Context, request *dto.ListPublishedVideoRequest) (*dto.ListPublishedVideoResponse, error)
}

var _ VideoUsecase = (*VideoUseCase)(nil)
