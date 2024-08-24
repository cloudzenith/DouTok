package videoapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/videooptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	base *baseadapter.Adapter
	core *svcoreadapter.Adapter
}

func New(
	base *baseadapter.Adapter,
	core *svcoreadapter.Adapter,
) *Application {
	return &Application{
		base: base,
		core: core,
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

	_, err = a.core.Feed(ctx, userId, options...)
	if err != nil {
		log.Context(ctx).Errorf("failed to feed short video: %v", err)
		return nil, errorx.New(1, "failed to feed short video")
	}

	// TODO: 视频上传调通后增加返回值内容
	return &svapi.FeedShortVideoResponse{}, nil
}

func (a *Application) GetVideoById(ctx context.Context, request *svapi.GetVideoByIdRequest) (*svapi.GetVideoByIdResponse, error) {
	_, err := a.core.GetVideoById(ctx, request.GetVideoId())
	if err != nil {
		log.Context(ctx).Errorf("failed to get video by id: %v", err)
		return nil, errorx.New(1, "failed to get video by id")
	}

	// TODO: 视频上传调通后增加返回值内容
	return &svapi.GetVideoByIdResponse{}, nil
}

func (a *Application) listPublishedList(ctx context.Context, userId int64, page, size int32) (*svapi.ListPublishedVideoResponse, error) {
	_, err := a.core.ListUserPublishedList(ctx, userId, page, size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list published video: %v", err)
		return nil, errorx.New(1, "failed to list published video")
	}

	// TODO: 上传视频调通后完善返回值
	return &svapi.ListPublishedVideoResponse{}, nil
}

func (a *Application) ListOthersPublishedVideo(ctx context.Context, request *svapi.ListOthersPublishedVideoRequest) (*svapi.ListPublishedVideoResponse, error) {
	return a.listPublishedList(ctx, request.AccountId, request.Pagination.Page, request.Pagination.Size)
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
	resp, err := a.base.PreSign4UploadVideo(
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

	// TODO: 通过sv-core预先添加视频基础信息
	videoId, err := a.core.PreSaveVideoInfo(ctx, request.Title)
	if err != nil {
		log.Context(ctx).Errorf("failed to pre save video info: %v", err)
		return nil, errorx.New(1, "failed to upload video")
	}

	return &svapi.PreSign4UploadVideoResponse{
		FileId:  resp.FileId,
		Url:     resp.Url,
		VideoId: videoId,
	}, nil
}

func (a *Application) ReportFinishUpload(ctx context.Context, request *svapi.ReportFinishUploadRequest) (*svapi.ListVideo, error) {
	if err := a.base.ReportUploaded(ctx, request.FileId); err != nil {
		log.Context(ctx).Errorf("failed to report finish upload: %v", err)
		return nil, errorx.New(1, "failed to report finish upload")
	}

	// TODO：通过sv-core标记视频已被上传成功，并得到视频基础信息

	return nil, nil
}

var _ svapi.ShortVideoCoreVideoServiceHTTPServer = (*Application)(nil)
