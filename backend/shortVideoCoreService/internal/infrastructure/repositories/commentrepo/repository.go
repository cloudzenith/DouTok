package commentrepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"gorm.io/gen"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) Create(ctx context.Context, comment *model.Comment) error { //TODO implement me
	return query.Q.WithContext(ctx).Comment.Create(comment)
}

func (r *PersistRepository) RemoveById(ctx context.Context, commentId int64) error {
	_, err := query.Q.WithContext(ctx).Comment.Where(query.Comment.ID.Eq(commentId)).Update(
		query.Comment.IsDeleted, true,
	)
	return err
}

func (r *PersistRepository) ListByVideoId(ctx context.Context, videoId int64, limit, offset int) ([]*model.Comment, error) {
	return query.Q.WithContext(ctx).Comment.Where(
		query.Q.Comment.VideoID.Eq(videoId),
	).Limit(limit).Offset(offset).Find()
}

func (r *PersistRepository) GetById(ctx context.Context, commentId int64) (*model.Comment, error) {
	return query.Q.WithContext(ctx).Comment.Where(query.Comment.ID.Eq(commentId)).First()
}

func (r *PersistRepository) GetByIdList(ctx context.Context, commentIdList []int64) ([]*model.Comment, error) {
	return query.Q.WithContext(ctx).Comment.Where(query.Comment.ID.In(commentIdList...)).Find()
}

func (r *PersistRepository) count(ctx context.Context, conditions ...gen.Condition) (int64, error) {
	return query.Q.WithContext(ctx).Comment.Where(conditions...).Count()
}

func (r *PersistRepository) CountByVideoId(ctx context.Context, videoId int64) (int64, error) {
	return r.count(ctx, query.Comment.VideoID.Eq(videoId))
}

func (r *PersistRepository) CountByUserId(ctx context.Context, userId int64) (int64, error) {
	return r.count(ctx, query.Comment.UserID.Eq(userId))
}

var _ repoiface.CommentRepository = (*PersistRepository)(nil)
