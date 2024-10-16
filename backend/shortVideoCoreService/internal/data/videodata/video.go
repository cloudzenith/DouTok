package videodata

import (
	"context"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"time"
)

type VideoRepo struct {
}

// NewVideoRepo .
func NewVideoRepo() *VideoRepo {
	return &VideoRepo{}
}

func (r *VideoRepo) Save(ctx context.Context, tx *query.Query, v *model.Video) error {
	return tx.Video.WithContext(ctx).Create(v)
}

func (r *VideoRepo) UpdateById(ctx context.Context, tx *query.Query, v *model.Video) (int64, error) {
	res, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(v.ID)).Updates(v)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}

func (r *VideoRepo) FindByID(ctx context.Context, tx *query.Query, id int64) (*model.Video, error) {
	video, err := tx.Video.WithContext(ctx).Where(tx.Video.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return video, nil
}

func (r *VideoRepo) FindByIdList(ctx context.Context, idList []int64) ([]*model.Video, error) {
	return query.Q.WithContext(ctx).Video.Where(query.Q.Video.ID.In(idList...)).Find()
}

func (r *VideoRepo) GetVideoList(
	ctx context.Context, tx *query.Query, userId int64, latestTime int64, pageRequest *infra_dto.PaginationRequest,
) ([]*model.Video, *infra_dto.PaginationResponse, error) {
	//var sortStrings []string
	//for _, sortField := range request.PaginationRequest.SortFields {
	//	order := "ASC"
	//	if sortField.Order == 1 {
	//		order = "DESC"
	//	}
	//	sortStrings = append(sortStrings, fmt.Sprintf("%s %s", sortField.Field, order))
	//}
	//sortString := strings.Join(sortStrings, ", ")
	// 暂时屏蔽外层排序逻辑
	offset := (pageRequest.PageNum - 1) * pageRequest.PageSize
	videos, err := tx.Video.WithContext(ctx).
		Where(tx.Video.UserID.Eq(userId)).
		Where(tx.Video.CreatedAt.Lt(time.Unix(latestTime, 0).UTC())).
		Limit(int(pageRequest.PageSize)).
		Offset(int(offset)).
		Order(tx.Video.ID.Desc()).Find()
	if err != nil {
		return nil, nil, err
	}
	return videos, &infra_dto.PaginationResponse{
		Page:  int64(pageRequest.PageNum),
		Count: int64(len(videos)),
	}, nil
}

func (r *VideoRepo) GetVideoFeed(ctx context.Context, tx *query.Query, userId, latestTime, num int64) ([]*model.Video, error) {
	videos, err := tx.Video.WithContext(ctx).
		Where(tx.Video.CreatedAt.Lt(time.Unix(latestTime, 0).UTC())).
		Limit(int(num)).
		Order(tx.Video.ID.Desc()).Find()
	if err != nil {
		return nil, err
	}
	return videos, nil
}
