package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/videooptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) ListUserPublishedList(ctx context.Context, userId int64, pageIndex, pageSize int32) (*v1.ListPublishedVideoResponse, error) {
	req := &v1.ListPublishedVideoRequest{
		UserId: userId,
		Pagination: &v1.PaginationRequest{
			Page: pageIndex,
			Size: pageSize,
		},
	}
	resp, err := a.video.ListPublishedVideo(ctx, req)
	return respcheck.CheckT[*v1.ListPublishedVideoResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListPublishedVideoResponse {
			return resp
		},
	)
}

func (a *Adapter) Feed(ctx context.Context, userId int64, num int64, options ...videooptions.FeedOptions) (*v1.FeedShortVideoResponse, error) {
	req := &v1.FeedShortVideoRequest{
		LatestTime: 0,
		UserId:     userId,
		FeedNum:    num,
	}
	for _, opt := range options {
		opt(req)
	}

	resp, err := a.video.FeedShortVideo(ctx, req)
	return respcheck.CheckT[*v1.FeedShortVideoResponse, *v1.Metadata](
		resp, err,
		func() *v1.FeedShortVideoResponse {
			return resp
		},
	)
}

func (a *Adapter) GetVideoById(ctx context.Context, videoId int64) (*v1.Video, error) {
	req := &v1.GetVideoByIdRequest{
		VideoId: videoId,
	}
	resp, err := a.video.GetVideoById(ctx, req)
	return respcheck.CheckT[*v1.Video, *v1.Metadata](
		resp, err,
		func() *v1.Video {
			return resp.GetVideo()
		},
	)
}

func (a *Adapter) SaveVideoInfo(ctx context.Context, title, videoUrl, coverUrl, desc string, userId int64) (int64, error) {
	req := &v1.PublishVideoRequest{
		Title:       title,
		PlayUrl:     videoUrl,
		CoverUrl:    coverUrl,
		Description: desc,
		UserId:      userId,
	}
	resp, err := a.video.PublishVideo(ctx, req)
	return respcheck.CheckT[int64, *v1.Metadata](
		resp, err,
		func() int64 {
			return resp.VideoId
		},
	)
}

func (a *Adapter) GetVideosByIdList(ctx context.Context, idList []int64) ([]*v1.Video, error) {
	req := &v1.GetVideoByIdListRequest{
		VideoIdList: idList,
	}
	resp, err := a.video.GetVideoByIdList(ctx, req)
	return respcheck.CheckT[[]*v1.Video, *v1.Metadata](
		resp, err,
		func() []*v1.Video {
			return resp.Videos
		},
	)
}
