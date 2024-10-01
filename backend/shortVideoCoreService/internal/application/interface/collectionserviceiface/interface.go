package collectionserviceiface

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/collection"
)

type CollectionService interface {
	CreateCollection(ctx context.Context, userId int64, name, description string) error
	GetCollectionById(ctx context.Context, collectionId int64) (*collection.Collection, error)
	RemoveCollection(ctx context.Context, collectionId int64) error
	ListCollection(ctx context.Context, userId int64, limit, offset int) ([]*collection.Collection, error)
	UpdateCollection(ctx context.Context, collectionId int64, name, description string) error
	AddVideo2Collection(ctx context.Context, collectionId, videoId int64) error
	RemoveVideo2Collection(ctx context.Context, collectionId, videoId int64) error
	ListCollectionVideo(ctx context.Context, collectionId int64, pagination *v1.PaginationRequest) ([]int64, error)
}
