package videodata

import (
	"context"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
)

type IVideoRepo interface {
	Save(ctx context.Context, tx *query.Query, v *model.Video) error
	UpdateById(ctx context.Context, tx *query.Query, v *model.Video) (int64, error)
	FindByID(ctx context.Context, tx *query.Query, id int64) (*model.Video, error)
	FindByIdList(ctx context.Context, idList []int64) ([]*model.Video, error)
	GetVideoList(
		ctx context.Context, tx *query.Query, userId int64, latestTime int64, PaginationRequest *infra_dto.PaginationRequest,
	) ([]*model.Video, *infra_dto.PaginationResponse, error)
	GetVideoFeed(ctx context.Context, tx *query.Query, userId, latestTime, num int64) ([]*model.Video, error)
}

var _ IVideoRepo = (*VideoRepo)(nil)
