package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *model.Comment) error
	RemoveById(ctx context.Context, commentId int64) error
	ListByVideoId(ctx context.Context, videoId int64, limit, offset int) ([]*model.Comment, error)
	GetById(ctx context.Context, commentId int64) (*model.Comment, error)
	GetByIdList(ctx context.Context, commentIdList []int64) ([]*model.Comment, error)
	CountByVideoId(ctx context.Context, videoId int64) (int64, error)
	CountByUserId(ctx context.Context, userId int64) (int64, error)
}
