package videodata

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/db"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type VideoRepo struct {
	dbClient *db.DBClient
	log      *log.Helper
}

// NewVideoRepo .
func NewVideoRepo(dbClient *db.DBClient, logger log.Logger) *VideoRepo {
	return &VideoRepo{
		dbClient: dbClient,
		log:      log.NewHelper(logger),
	}
}

func (r *VideoRepo) Save(ctx context.Context, v *model.Video) error {
	result := r.dbClient.DB(ctx).Create(v)
	return result.Error
}

func (r *VideoRepo) UpdateById(ctx context.Context, v *model.Video) (int64, error) {
	result := r.dbClient.DB(ctx).Where(&model.Video{ID: v.ID}).Updates(v)
	return result.RowsAffected, result.Error
}

func (r *VideoRepo) FindByID(ctx context.Context, id int64) (*model.Video, error) {
	video := &model.Video{}
	result := r.dbClient.DB(ctx).First(video, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return video, nil
}

func (r *VideoRepo) GetVideoList(ctx context.Context, request *dto.GetVideoListRequest) (*dto.GetVideoListResponse, error) {
	whereConditionFn := func(db *gorm.DB) *gorm.DB {
		db = db.
			Where("user_id = ?", request.UserId).
			Where("created_at < ?", time.Unix(request.LatestTime, 0).UTC())
		return db
	}

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
	sortString := "id desc"

	var videos []*model.Video
	result := r.dbClient.WhereWithPaginateAndSort(ctx, whereConditionFn, &videos, sortString, request.PaginationRequest)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.GetVideoListResponse{
		Videos: videos,
		PaginationResponse: &infra_dto.PaginationResponse{
			Page:  int64(request.PaginationRequest.PageNum),
			Count: int64(len(videos)),
		},
	}, nil
}

func (r *VideoRepo) GetVideoFeed(ctx context.Context, request *dto.GetVideoFeedRequest) (*dto.GetVideoFeedResponse, error) {
	var videos []*model.Video
	result := r.dbClient.DB(ctx).
		Where("user_id = ?", request.UserId).
		Where("created_at < ?", time.Unix(request.LatestTime, 0).UTC()).
		Limit(int(request.Num)).Order("id desc").Find(&videos)
	if result.Error != nil {
		return nil, result.Error
	}
	return &dto.GetVideoFeedResponse{
		Videos: videos,
	}, nil
}
