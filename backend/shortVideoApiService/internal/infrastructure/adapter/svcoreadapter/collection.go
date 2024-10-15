package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/v2/log"
)

func (a *Adapter) AddVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) error {
	req := &v1.AddVideo2CollectionRequest{
		CollectionId: collectionId,
		VideoId:      videoId,
		UserId:       userId,
	}

	resp, err := a.collection.AddVideo2Collection(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) RemoveVideoFromCollection(ctx context.Context, userId, collectionId, videoId int64) error {
	req := &v1.RemoveVideoFromCollectionRequest{
		CollectionId: collectionId,
		VideoId:      videoId,
		UserId:       userId,
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

func (a *Adapter) IsCollected(ctx context.Context, userId int64, videoIdList []int64) (map[int64]bool, error) {
	req := &v1.IsCollectedRequest{
		UserId:      userId,
		VideoIdList: videoIdList,
	}

	resp, err := a.collection.IsCollected(ctx, req)
	log.Context(ctx).Infof("IsCollected resp: %v", resp)
	return respcheck.CheckT[map[int64]bool, *v1.Metadata](
		resp, err,
		func() map[int64]bool {
			result := make(map[int64]bool)
			if len(resp.VideoIdList) == 0 {
				return result
			}

			for _, item := range resp.VideoIdList {
				result[item] = true
			}

			return result
		},
	)
}

func (a *Adapter) CountCollected4Video(ctx context.Context, videoIdList []int64) (map[int64]int64, error) {
	req := &v1.CountCollect4VideoRequest{
		VideoIdList: videoIdList,
	}

	resp, err := a.collection.CountCollect4Video(ctx, req)
	return respcheck.CheckT[map[int64]int64, *v1.Metadata](
		resp, err,
		func() map[int64]int64 {
			result := make(map[int64]int64)
			for _, item := range resp.CountResult {
				result[item.Id] = item.Count
			}

			return result
		},
	)
}
