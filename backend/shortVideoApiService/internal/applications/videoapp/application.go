package videoapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/interface/videoserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/videooptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	base         *baseadapter.Adapter
	core         *svcoreadapter.Adapter
	videoService videoserviceiface.VideoService
}

func New(
	base *baseadapter.Adapter,
	core *svcoreadapter.Adapter,
	videoService videoserviceiface.VideoService,
) *Application {
	return &Application{
		base:         base,
		core:         core,
		videoService: videoService,
	}
}

func (a *Application) FeedShortVideo(ctx context.Context, request *svapi.FeedShortVideoRequest) (*svapi.FeedShortVideoResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user id from context: %v", err)
		return nil, errorx.New(1, "failed to get user id from context")
	}

	var options []videooptions.FeedOptions
	if request.LatestTime != 0 {
		options = append(options, videooptions.FeedWithLatestTime(request.LatestTime))
	}

	resp, err := a.core.Feed(ctx, userId, request.FeedNum, options...)
	if err != nil {
		log.Context(ctx).Errorf("failed to feed short video: %v", err)
		return nil, errorx.New(1, "failed to feed short video")
	}

	videos := dto.ToPBVideoList(resp.Videos)
	a.videoService.AssembleUserIsFollowing(ctx, videos, userId)
	a.videoService.AssembleVideoCountInfo(ctx, videos)

	return &svapi.FeedShortVideoResponse{
		Videos: videos,
	}, nil
}

func (a *Application) GetVideoById(ctx context.Context, request *svapi.GetVideoByIdRequest) (*svapi.GetVideoByIdResponse, error) {
	video, err := a.core.GetVideoById(ctx, request.GetVideoId())
	if err != nil {
		log.Context(ctx).Errorf("failed to get video by id: %v", err)
		return nil, errorx.New(1, "failed to get video by id")
	}

	return &svapi.GetVideoByIdResponse{
		Video: dto.ToPBVideo(video),
	}, nil
}

func (a *Application) listPublishedList(ctx context.Context, userId int64, page, size int32) (*svapi.ListPublishedVideoResponse, error) {
	resp, err := a.core.ListUserPublishedList(ctx, userId, page, size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list published video: %v", err)
		return nil, errorx.New(1, "failed to list published video")
	}

	videoList, err := a.videoService.AssembleVideoList(ctx, userId, resp.Videos)
	if err != nil {
		log.Context(ctx).Errorf("failed to assemble video list: %v", err)
		return nil, errorx.New(1, "failed to assemble video list")
	}

	a.videoService.AssembleUserIsFollowing(ctx, videoList, userId)
	a.videoService.AssembleVideoCountInfo(ctx, videoList)

	return &svapi.ListPublishedVideoResponse{
		VideoList: videoList,
		Pagination: &svapi.PaginationResponse{
			Page:  resp.Pagination.Page,
			Total: resp.Pagination.Total,
			Count: resp.Pagination.Count,
		},
	}, nil
}

func (a *Application) ListPublishedVideo(ctx context.Context, request *svapi.ListPublishedVideoRequest) (*svapi.ListPublishedVideoResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user id from context: %v", err)
		return nil, errorx.New(1, "failed to get user id from context")
	}

	return a.listPublishedList(ctx, userId, request.Pagination.Page, request.Pagination.Size)
}

func (a *Application) PreSign4UploadVideo(ctx context.Context, request *svapi.PreSign4UploadVideoRequest) (*svapi.PreSign4UploadVideoResponse, error) {
	resp, err := a.base.PreSign4Upload(
		ctx,
		request.Hash,
		request.FileType,
		request.Filename,
		request.Size,
		3600,
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to pre sign for upload video: %v", err)
		return nil, errorx.New(1, "failed to pre sign for upload video")
	}

	return &svapi.PreSign4UploadVideoResponse{
		FileId: resp.FileId,
		Url:    resp.Url,
	}, nil
}

func (a *Application) ReportVideoFinishUpload(ctx context.Context, request *svapi.ReportVideoFinishUploadRequest) (*svapi.ReportVideoFinishUploadResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		log.Context(ctx).Errorf("failed to get user id from context: %v", err)
		return nil, errorx.New(1, "failed to get user id from context")
	}

	_, err = a.base.ReportPublicUploaded(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to report finish upload: %v", err)
		return nil, errorx.New(1, "failed to report finish upload")
	}

	videoId, err := a.core.SaveVideoInfo(ctx, request.Title, request.VideoUrl, request.CoverUrl, request.Description, userId)
	if err != nil {
		log.Context(ctx).Errorf("failed to save video info: %v", err)
		return nil, errorx.New(1, "failed to save video info")
	}

	return &svapi.ReportVideoFinishUploadResponse{
		VideoId: videoId,
	}, nil
}

func (a *Application) ReportFinishUpload(ctx context.Context, request *svapi.ReportFinishUploadRequest) (*svapi.ReportFinishUploadResponse, error) {
	resp, err := a.base.ReportUploaded(ctx, request.FileId)
	if err != nil {
		log.Context(ctx).Errorf("failed to report finish upload: %v", err)
		return nil, errorx.New(1, "failed to report finish upload")
	}

	return &svapi.ReportFinishUploadResponse{
		Url: resp.Url,
	}, nil
}

func (a *Application) PreSign4UploadCover(ctx context.Context, request *svapi.PreSign4UploadRequest) (*svapi.PreSign4UploadResponse, error) {
	resp, err := a.base.PreSign4Upload(
		ctx,
		request.Hash,
		request.FileType,
		request.Filename,
		request.Size,
		3600,
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to pre sign for upload cover: %v", err)
		return nil, errorx.New(1, "failed to pre sign for upload cover")
	}

	return &svapi.PreSign4UploadResponse{
		FileId: resp.FileId,
		Url:    resp.Url,
	}, nil
}

var _ svapi.ShortVideoCoreVideoServiceHTTPServer = (*Application)(nil)
