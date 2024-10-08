package videoapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/baseadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/videooptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
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

	resp, err := a.core.Feed(ctx, userId, request.FeedNum, options...)
	if err != nil {
		log.Context(ctx).Errorf("failed to feed short video: %v", err)
		return nil, errorx.New(1, "failed to feed short video")
	}

	return &svapi.FeedShortVideoResponse{
		Videos: dto.ToPBVideoList(resp.Videos),
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

func (a *Application) assembleVideoList(ctx context.Context, userId int64, data []*v1.Video) ([]*svapi.Video, error) {
	var result []*svapi.Video

	var videoIdList []int64
	for _, video := range data {
		videoIdList = append(videoIdList, video.GetId())
	}

	isFavoriteMap, err := a.core.IsUserFavoriteVideo(ctx, userId, videoIdList)
	if err != nil {
		// 弱依赖
		log.Context(ctx).Warnf("failed to check favorite video: %v", err)
	}

	for _, video := range data {
		isFavorite, ok := isFavoriteMap[video.GetId()]

		result = append(result, &svapi.Video{
			Id:            video.GetId(),
			Title:         video.Title,
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite && ok,
			Author: &svapi.VideoAuthor{
				Id:          video.Author.Id,
				Name:        video.Author.Name,
				Avatar:      video.Author.Avatar,
				IsFollowing: video.Author.IsFollowing != 0,
			},
		})
	}

	return result, nil
}

func (a *Application) listPublishedList(ctx context.Context, userId int64, page, size int32) (*svapi.ListPublishedVideoResponse, error) {
	resp, err := a.core.ListUserPublishedList(ctx, userId, page, size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list published video: %v", err)
		return nil, errorx.New(1, "failed to list published video")
	}

	videoList, err := a.assembleVideoList(ctx, userId, resp.Videos)
	if err != nil {
		log.Context(ctx).Errorf("failed to assemble video list: %v", err)
		return nil, errorx.New(1, "failed to assemble video list")
	}

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
