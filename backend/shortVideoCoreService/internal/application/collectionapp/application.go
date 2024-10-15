package collectionapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/collectionserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	collection collectionserviceiface.CollectionService
	v1.UnimplementedCollectionServiceServer
}

func New(collection collectionserviceiface.CollectionService) *Application {
	return &Application{
		collection: collection,
	}
}

func (a *Application) CreateCollection(ctx context.Context, request *v1.CreateCollectionRequest) (*v1.CreateCollectionResponse, error) {
	err := a.collection.CreateCollection(ctx, request.UserId, request.Name, request.Description)
	if err != nil {
		log.Context(ctx).Errorf("CreateCollection error: %v", err)
		return &v1.CreateCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CreateCollectionResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) GetCollectionById(ctx context.Context, request *v1.GetCollectionByIdRequest) (*v1.GetCollectionByIdResponse, error) {
	data, err := a.collection.GetCollectionById(ctx, request.GetId())
	if err != nil {
		log.Context(ctx).Errorf("GetCollectionById error: %v", err)
		return &v1.GetCollectionByIdResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.GetCollectionByIdResponse{
		Meta: utils.GetSuccessMeta(),
		Collection: &v1.Collection{
			Id:          data.ID,
			UserId:      data.UserId,
			Name:        data.Title,
			Description: data.Description,
		},
	}, nil
}

func (a *Application) RemoveCollection(ctx context.Context, request *v1.RemoveCollectionRequest) (*v1.RemoveCollectionResponse, error) {
	if err := a.collection.RemoveCollection(ctx, request.GetId()); err != nil {
		log.Context(ctx).Errorf("RemoveCollection error: %v", err)
		return &v1.RemoveCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.RemoveCollectionResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) ListCollection(ctx context.Context, request *v1.ListCollectionRequest) (*v1.ListCollectionResponse, error) {
	limit, offset := utils.GetLimitOffset(
		int(request.GetPagination().GetPage()),
		int(request.GetPagination().GetSize()),
	)

	data, err := a.collection.ListCollection(
		ctx,
		request.GetUserId(),
		limit, offset,
	)
	if err != nil {
		log.Context(ctx).Errorf("ListCollection error: %v", err)
		return &v1.ListCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var collections []*v1.Collection
	for _, c := range data.Data {
		collections = append(collections, &v1.Collection{
			Id:          c.ID,
			UserId:      c.UserId,
			Name:        c.Title,
			Description: c.Description,
		})
	}

	return &v1.ListCollectionResponse{
		Meta:        utils.GetSuccessMeta(),
		Collections: collections,
		Pagination:  utils.GetPageResponse(data.Count, request.Pagination.Page, request.Pagination.Size),
	}, nil
}

func (a *Application) UpdateCollection(ctx context.Context, request *v1.UpdateCollectionRequest) (*v1.UpdateCollectionResponse, error) {
	err := a.collection.UpdateCollection(ctx, request.GetId(), request.Name, request.Description)
	if err != nil {
		log.Context(ctx).Errorf("UpdateCollection error: %v", err)
		return &v1.UpdateCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.UpdateCollectionResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) AddVideo2Collection(ctx context.Context, request *v1.AddVideo2CollectionRequest) (*v1.AddVideo2CollectionResponse, error) {
	err := a.collection.GenerateDefaultCollection(ctx, request.GetUserId())
	if err != nil {
		log.Context(ctx).Errorf("failed to check default collection: %v", err)
		return &v1.AddVideo2CollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	err = a.collection.AddVideo2Collection(ctx, request.GetUserId(), request.GetCollectionId(), request.GetVideoId())
	if err != nil {
		log.Context(ctx).Errorf("AddVideo2Collection error: %v", err)
		return &v1.AddVideo2CollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.AddVideo2CollectionResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) RemoveVideoFromCollection(ctx context.Context, request *v1.RemoveVideoFromCollectionRequest) (*v1.RemoveVideoFromCollectionResponse, error) {
	err := a.collection.GenerateDefaultCollection(ctx, request.GetUserId())
	if err != nil {
		log.Context(ctx).Errorf("failed to check default collection: %v", err)
		return &v1.RemoveVideoFromCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	err = a.collection.RemoveVideo2Collection(ctx, request.GetUserId(), request.GetCollectionId(), request.GetVideoId())
	if err != nil {
		log.Context(ctx).Errorf("RemoveVideoFromCollection error: %v", err)
		return &v1.RemoveVideoFromCollectionResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.RemoveVideoFromCollectionResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) ListCollectionVideo(ctx context.Context, request *v1.ListCollectionVideoRequest) (*v1.ListCollectionVideoResponse, error) {
	data, err := a.collection.ListCollectionVideo(ctx, request.GetCollectionId(), request.GetPagination())
	if err != nil {
		log.Context(ctx).Errorf("ListCollectionVideo error: %v", err)
		return &v1.ListCollectionVideoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.ListCollectionVideoResponse{
		Meta:        utils.GetSuccessMeta(),
		VideoIdList: data.Data,
		Pagination:  utils.GetPageResponse(data.Count, request.Pagination.Page, request.Pagination.Size),
	}, nil
}

func (a *Application) IsCollected(ctx context.Context, request *v1.IsCollectedRequest) (*v1.IsCollectedResponse, error) {
	data, err := a.collection.ListCollectedVideoByGiven(ctx, request.GetUserId(), request.GetVideoIdList())
	if err != nil {
		log.Context(ctx).Errorf("IsCollected error: %v", err)
		return &v1.IsCollectedResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.IsCollectedResponse{
		Meta:        utils.GetSuccessMeta(),
		VideoIdList: data,
	}, nil
}

func (a *Application) CountCollect4Video(ctx context.Context, request *v1.CountCollect4VideoRequest) (*v1.CountCollect4VideoResponse, error) {
	countInfo, err := a.collection.CountCollectedNumber4Video(ctx, request.GetVideoIdList())
	if err != nil {
		log.Context(ctx).Errorf("CountCollect4Video error: %v", err)
		return &v1.CountCollect4VideoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var results []*v1.CountCollect4VideoResult
	for _, item := range countInfo {
		results = append(results, &v1.CountCollect4VideoResult{
			Id:    item.Id,
			Count: item.Count,
		})
	}

	return &v1.CountCollect4VideoResponse{
		Meta:        utils.GetSuccessMeta(),
		CountResult: results,
	}, nil
}
