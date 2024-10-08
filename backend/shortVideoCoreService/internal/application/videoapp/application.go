package videoapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	service_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/dto"
	domain_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/videodomain"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type VideoApplication struct {
	videoUsecase videodomain.VideoUsecase
	v1.UnimplementedVideoServiceServer
}

func NewVideoApplication(video videodomain.VideoUsecase) *VideoApplication {
	return &VideoApplication{
		videoUsecase: video,
	}
}

func (s *VideoApplication) PublishVideo(ctx context.Context, in *v1.PublishVideoRequest) (*v1.PublishVideoResponse, error) {
	videoId, err := s.videoUsecase.PublishVideo(ctx, &domain_dto.PublishVideoRequest{
		UserId:      in.UserId,
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
			BizCode: 0,
			Message: "success",
		},
		VideoId: videoId,
	}, nil
}

func (s *VideoApplication) GetVideoById(ctx context.Context, in *v1.GetVideoByIdRequest) (*v1.GetVideoByIdResponse, error) {
	video, err := s.videoUsecase.GetVideoById(ctx, in.VideoId)
	if err != nil {
		return nil, err
	}
	return &v1.GetVideoByIdResponse{
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
		Video: service_dto.ToPBVideo(video),
	}, nil
}

func (s *VideoApplication) FeedShortVideo(ctx context.Context, in *v1.FeedShortVideoRequest) (*v1.FeedShortVideoResponse, error) {
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
			BizCode: 0,
			Message: "success",
		},
		Videos: service_dto.ToPBVideoList(resp.Videos),
	}, nil
}

func (s *VideoApplication) ListPublishedVideo(ctx context.Context, in *v1.ListPublishedVideoRequest) (*v1.ListPublishedVideoResponse, error) {
	resp, err := s.videoUsecase.ListPublishedVideo(ctx, &domain_dto.ListPublishedVideoRequest{
		UserId:     in.UserId,
		Pagination: infra_dto.FromPBPaginationRequest(in.Pagination),
	})
	if err != nil {
		return nil, err
	}
	return &v1.ListPublishedVideoResponse{
		Meta: &v1.Metadata{
			BizCode: 0,
			Message: "success",
		},
		Videos:     service_dto.ToPBVideoList(resp.Videos),
		Pagination: infra_dto.ToPBPaginationResponse(resp.Pagination),
	}, nil
}

func (s *VideoApplication) GetVideoByIdList(ctx context.Context, in *v1.GetVideoByIdListRequest) (*v1.GetVideoByIdListResponse, error) {
	data, err := s.videoUsecase.GetVideoByIdList(ctx, in.VideoIdList)
	if err != nil {
		log.Context(ctx).Errorf("GetVideoByIdList error: %v", err)
		return &v1.GetVideoByIdListResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var result []*v1.Video
	for _, item := range data {
		result = append(result, item.ToPB())
	}

	return &v1.GetVideoByIdListResponse{
		Meta:   utils.GetSuccessMeta(),
		Videos: result,
	}, nil
}
