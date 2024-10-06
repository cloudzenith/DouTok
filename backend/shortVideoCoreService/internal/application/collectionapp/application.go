package collectionapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/collectionserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
	"math"
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
	data, err := a.collection.ListCollection(
		ctx,
		request.GetUserId(),
		int(request.GetPagination().GetSize()),
		int(request.GetPagination().GetPage()),
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
		Pagination: &v1.PaginationResponse{
			Page:  request.Pagination.Page,
			Total: int32(math.Ceil(float64(data.Count) / float64(request.Pagination.Size))),
			Count: int32(data.Count),
		},
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
	err := a.collection.AddVideo2Collection(ctx, request.GetCollectionId(), request.GetVideoId())
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
	err := a.collection.RemoveVideo2Collection(ctx, request.GetCollectionId(), request.GetVideoId())
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
		Pagination: &v1.PaginationResponse{
			Page:  request.Pagination.Page,
			Total: int32(math.Ceil(float64(data.Count) / float64(request.Pagination.Size))),
			Count: int32(data.Count),
		},
	}, nil
}
