package collectionserviceiface

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/collection"
)

type ListCollectionResult struct {
	Data  []*collection.Collection
	Count int64
}

type ListCollectionVideoResult struct {
	Data  []int64
	Count int64
}

type CountResult struct {
	Id    int64
	Count int64
}

type CollectionService interface {
	CreateCollection(ctx context.Context, userId int64, name, description string) error
	GetCollectionById(ctx context.Context, collectionId int64) (*collection.Collection, error)
	RemoveCollection(ctx context.Context, collectionId int64) error
	ListCollection(ctx context.Context, userId int64, limit, offset int) (*ListCollectionResult, error)
	UpdateCollection(ctx context.Context, collectionId int64, name, description string) error
	AddVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) error
	RemoveVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) error
	ListCollectionVideo(ctx context.Context, collectionId int64, pagination *v1.PaginationRequest) (*ListCollectionVideoResult, error)
	ListCollectedVideoByGiven(ctx context.Context, userId int64, videoIdList []int64) ([]int64, error)
	GenerateDefaultCollection(ctx context.Context, userId int64) error
	CountCollectedNumber4Video(ctx context.Context, videoId []int64) ([]*CountResult, error)
}
