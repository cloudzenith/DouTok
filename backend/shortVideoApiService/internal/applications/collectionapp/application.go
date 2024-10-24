package collectionapp

import (
	"context"
	"errors"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/applications/interface/videoserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/claims"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/errorx"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	core         *svcoreadapter.Adapter
	videoService videoserviceiface.VideoService
}

func New(
	core *svcoreadapter.Adapter,
	videoService videoserviceiface.VideoService,
) *Application {
	return &Application{
		core:         core,
		videoService: videoService,
	}
}

func (a *Application) checkCollectionBelongUser(ctx context.Context, collectionId int64) error {
	if collectionId == 0 {
		log.Context(ctx).Warnf("collectionId is empty")
		return nil
	}

	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return errorx.New(1, "获取用户信息失败")
	}

	data, err := a.core.GetCollectionById(ctx, collectionId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get collection info: %v", err)
		return errorx.New(1, "信息不存在")
	}

	if data.UserId != userId {
		return errors.New("此收藏夹不属于当前用户")
	}

	return nil
}

func (a *Application) AddVideo2Collection(ctx context.Context, request *svapi.AddVideo2CollectionRequest) (*svapi.AddVideo2CollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.core.AddVideo2Collection(ctx, userId, request.CollectionId, request.VideoId); err != nil {
		log.Context(ctx).Errorf("failed to add video to collection: %v", err)
		return nil, errorx.New(1, "添加失败")
	}

	return &svapi.AddVideo2CollectionResponse{}, nil
}

func (a *Application) CreateCollection(ctx context.Context, request *svapi.CreateCollectionRequest) (*svapi.CreateCollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.core.AddCollection(ctx, request.Name, request.Description, userId); err != nil {
		log.Context(ctx).Errorf("failed to create collection: %v", err)
		return nil, errorx.New(1, "创建失败")
	}

	return &svapi.CreateCollectionResponse{}, nil
}

func (a *Application) ListCollection(ctx context.Context, request *svapi.ListCollectionRequest) (*svapi.ListCollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	data, err := a.core.ListCollection(ctx, userId, request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		log.Context(ctx).Errorf("failed to list collection: %v", err)
		return nil, errorx.New(1, "获取失败")
	}

	var result []*svapi.Collection
	for _, item := range data.Collections {
		result = append(result, &svapi.Collection{
			Id:          item.Id,
			Name:        item.Name,
			Description: item.Description,
		})
	}

	return &svapi.ListCollectionResponse{
		Collections: result,
		Pagination: &svapi.PaginationResponse{
			Page:  data.Pagination.Page,
			Total: data.Pagination.Total,
			Count: data.Pagination.Count,
		},
	}, nil
}

func (a *Application) ListVideo4Collection(ctx context.Context, request *svapi.ListVideo4CollectionRequest) (*svapi.ListVideo4CollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	resp, err := a.core.ListVideo4Collection(ctx, request.CollectionId, request.Pagination.Page, request.Pagination.Size)
	if err != nil {
		return nil, errorx.New(1, "获取失败")
	}

	videoInfoList, err := a.core.GetVideosByIdList(ctx, resp.VideoIdList)
	if err != nil {
		log.Context(ctx).Errorf("failed to get video info: %v", err)
		return nil, errorx.New(1, "获取视频信息失败")
	}

	result, err := a.videoService.AssembleVideo(ctx, userId, videoInfoList)
	if err != nil {
		log.Context(ctx).Warnf("something wrong in assembling videos: %v", err)
	}

	return &svapi.ListVideo4CollectionResponse{
		Videos: result,
		Pagination: &svapi.PaginationResponse{
			Page:  resp.Pagination.Page,
			Total: resp.Pagination.Total,
			Count: resp.Pagination.Count,
		},
	}, nil
}

func (a *Application) RemoveCollection(ctx context.Context, request *svapi.RemoveCollectionRequest) (*svapi.RemoveCollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.Id); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.core.RemoveCollection(ctx, request.Id); err != nil {
		log.Context(ctx).Errorf("failed to remove collection: %v", err)
		return nil, errorx.New(1, "删除失败")
	}

	return &svapi.RemoveCollectionResponse{}, nil
}

func (a *Application) RemoveVideoFromCollection(ctx context.Context, request *svapi.RemoveVideoFromCollectionRequest) (*svapi.RemoveVideoFromCollectionResponse, error) {
	userId, err := claims.GetUserId(ctx)
	if err != nil {
		return nil, errorx.New(1, "获取用户信息失败")
	}

	if err := a.checkCollectionBelongUser(ctx, request.CollectionId); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.core.RemoveVideoFromCollection(ctx, userId, request.CollectionId, request.VideoId); err != nil {
		log.Context(ctx).Errorf("failed to remove video from collection: %v", err)
		return nil, errorx.New(1, "删除失败")
	}

	return &svapi.RemoveVideoFromCollectionResponse{}, nil
}

func (a *Application) UpdateCollection(ctx context.Context, request *svapi.UpdateCollectionRequest) (*svapi.UpdateCollectionResponse, error) {
	if err := a.checkCollectionBelongUser(ctx, request.Id); err != nil {
		return nil, errorx.New(1, err.Error())
	}

	if err := a.core.UpdateCollection(ctx, request.Id, request.Name, request.Description); err != nil {
		log.Context(ctx).Errorf("failed to update collection: %v", err)
		return nil, errorx.New(1, "更新失败")
	}

	return &svapi.UpdateCollectionResponse{}, nil
}

var _ svapi.CollectionServiceHTTPServer = (*Application)(nil)
