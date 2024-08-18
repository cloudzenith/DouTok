package videoservice

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	domain_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/videodomain"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	service_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/service/dto"
	"github.com/go-kratos/kratos/v2/log"
)

type VideoService struct {
	config       *conf.Config
	log          *log.Helper
	videoUsecase videodomain.VideoUsecase
	v1.UnimplementedVideoServiceServer
}

func NewVideoService(config *conf.Config, logger log.Logger, video videodomain.VideoUsecase) *VideoService {
	return &VideoService{
		config:       config,
		log:          log.NewHelper(logger),
		videoUsecase: video,
	}
}

func (s *VideoService) PublishVideo(ctx context.Context, in *v1.PublishVideoRequest) (*v1.PublishVideoResponse, error) {
	videoId, err := s.videoUsecase.PublishVideo(ctx, &domain_dto.PublishVideoRequest{
		Title:       in.Title,
		Description: in.Description,
		VideoURL:    in.PlayUrl,
		CoverURL:    in.CoverUrl,
	})
	if err != nil {
		return nil, err
	}
	return &v1.PublishVideoResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		VideoId: videoId,
	}, nil
}

func (s *VideoService) GetVideoById(ctx context.Context, in *v1.GetVideoByIdRequest) (*v1.GetVideoByIdResponse, error) {
	video, err := s.videoUsecase.GetVideoById(ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	return &v1.GetVideoByIdResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		Video: service_dto.ToPBVideo(video),
	}, nil
}

func (s *VideoService) FeedShortVideo(ctx context.Context, in *v1.FeedShortVideoRequest) (*v1.FeedShortVideoResponse, error) {
	resp, err := s.videoUsecase.FeedShortVideo(ctx, &domain_dto.FeedShortVideoRequest{
		UserId:     in.UserId,
		LatestTime: in.LatestTime,
		FeedNum:    in.FeedNum,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FeedShortVideoResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		Videos: service_dto.ToPBVideoList(resp.Videos),
	}, nil
}

func (s *VideoService) ListPublishedVideo(ctx context.Context, in *v1.ListPublishedVideoRequest) (*v1.ListPublishedVideoResponse, error) {
	resp, err := s.videoUsecase.ListPublishedVideo(ctx, &domain_dto.ListPublishedVideoRequest{
		UserId:     in.UserId,
		Pagination: infra_dto.FromPBPaginationRequest(in.Pagination),
	})
	if err != nil {
		return nil, err
	}
	return &v1.ListPublishedVideoResponse{
		Meta: &v1.Metadata{
			BizCode: 200,
			Message: "success",
		},
		Videos: service_dto.ToPBVideoList(resp.Videos),
	}, nil
}
