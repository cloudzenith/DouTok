package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/collectionserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type CollectionRepository interface {
	Create(ctx context.Context, collection *model.Collection) error
	GetById(ctx context.Context, id int64) (*model.Collection, error)
	RemoveById(ctx context.Context, id int64) error
	ListByUserId(ctx context.Context, userId int64, limit, offset int) ([]*model.Collection, error)
	ListFirstCollection4UserId(ctx context.Context, userId int64) (*model.Collection, error)
	CountByUserId(ctx context.Context, userId int64) (int64, error)
	Update(ctx context.Context, collection *model.Collection) error
	ListCollectionVideo(ctx context.Context, collectionId int64, limit, offset int) ([]*model.CollectionVideo, error)
	AddVideo2Collection(ctx context.Context, userId, collectionId, videoId int64) error
	RemoveVideoFromCollection(ctx context.Context, collectionId, videoId int64) error
	UpdateCollectionVideoTx(ctx context.Context, collectionVideo *model.CollectionVideo) error
	CountCollectionVideo(ctx context.Context, collectionId int64) (int64, error)
	ListCollectedVideoByGiven(ctx context.Context, userId int64, videoIdList []int64) ([]int64, error)
	GetCollectionVideo(ctx context.Context, collectionId, videoId int64) (*model.CollectionVideo, error)
	GetByIdTx(ctx context.Context, id int64) (*model.Collection, error)
	CountByVideoIdList(ctx context.Context, videoIdList []int64) ([]*collectionserviceiface.CountResult, error)
}
