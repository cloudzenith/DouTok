package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) AddVideo2Collection(ctx context.Context, collectionId, videoId int64) error {
	req := &v1.AddVideo2CollectionRequest{
		CollectionId: collectionId,
		VideoId:      videoId,
	}

	resp, err := a.collection.AddVideo2Collection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveVideoFromCollection(ctx context.Context, collectionId, videoId int64) error {
	req := &v1.RemoveVideoFromCollectionRequest{
		CollectionId: collectionId,
		VideoId:      videoId,
	}

	resp, err := a.collection.RemoveVideoFromCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) AddCollection(ctx context.Context, name, description string, userId int64) error {
	req := &v1.CreateCollectionRequest{
		Name:        name,
		Description: description,
		UserId:      userId,
	}

	resp, err := a.collection.CreateCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) ListCollection(ctx context.Context, userId int64, page, size int32) (*v1.ListCollectionResponse, error) {
	req := &v1.ListCollectionRequest{
		UserId: userId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.collection.ListCollection(ctx, req)
	return respcheck.CheckT[*v1.ListCollectionResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListCollectionResponse {
			return resp
		},
	)
}

func (a *Adapter) ListVideo4Collection(ctx context.Context, collectionId int64, page, size int32) (*v1.ListCollectionVideoResponse, error) {
	req := &v1.ListCollectionVideoRequest{
		CollectionId: collectionId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.collection.ListCollectionVideo(ctx, req)
	return respcheck.CheckT[*v1.ListCollectionVideoResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListCollectionVideoResponse {
			return resp
		},
	)
}

func (a *Adapter) RemoveCollection(ctx context.Context, collectionId int64) error {
	req := &v1.RemoveCollectionRequest{
		Id: collectionId,
	}

	resp, err := a.collection.RemoveCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) UpdateCollection(ctx context.Context, collectionId int64, name, description string) error {
	req := &v1.UpdateCollectionRequest{
		Id:          collectionId,
		Name:        name,
		Description: description,
	}

	resp, err := a.collection.UpdateCollection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) GetCollectionById(ctx context.Context, collectionId int64) (*v1.Collection, error) {
	req := &v1.GetCollectionByIdRequest{
		Id: collectionId,
	}

	resp, err := a.collection.GetCollectionById(ctx, req)
	return respcheck.CheckT[*v1.Collection, *v1.Metadata](
		resp, err,
		func() *v1.Collection {
			return resp.Collection
		},
	)
}
