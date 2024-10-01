package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type CollectionRepository interface {
	Create(ctx context.Context, collection *model.Collection) error
	GetById(ctx context.Context, id int64) (*model.Collection, error)
	RemoveById(ctx context.Context, id int64) error
	ListByUserId(ctx context.Context, userId int64, limit, offset int) ([]*model.Collection, error)
	Update(ctx context.Context, collection *model.Collection) error
	AddVideo2Collection(ctx context.Context, collectionId, videoId int64) error
	RemoveVideoFromCollection(ctx context.Context, collectionId, videoId int64) error
}
