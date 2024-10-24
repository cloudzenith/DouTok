package commentrepo

import (
	"context"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/commentserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/query"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (r *PersistRepository) Create(ctx context.Context, comment *model.Comment) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		return query.Q.WithContext(ctx).Comment.Create(comment)
	})
}

func (r *PersistRepository) Update(ctx context.Context, comment *model.Comment) error {
	return dbtx.TxDo(ctx, func(tx *query.QueryTx) error {
		_, err := query.Q.WithContext(ctx).Comment.Updates(comment)
		return err
	})
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

func (r *PersistRepository) listComment(ctx context.Context, conditions []gen.Condition, order []field.Expr, limit, offset int) ([]*model.Comment, error) {
	return query.Q.WithContext(ctx).Comment.Where(conditions...).Order(order...).Limit(limit).Offset(offset).Find()
}

func (r *PersistRepository) ListParentCommentByVideoId(ctx context.Context, videoId int64, limit, offset int) ([]*model.Comment, error) {
	return r.listComment(
		ctx,
		[]gen.Condition{
			query.Q.Comment.VideoID.Eq(videoId),
			query.Q.Comment.ParentID.Eq(0),
		},
		[]field.Expr{
			query.Q.Comment.ID.Asc(),
		},
		limit,
		offset,
	)
}

func (r *PersistRepository) ListChildCommentByCommentId(ctx context.Context, commentId int64, limit, offset int) ([]*model.Comment, error) {
	return r.listComment(
		ctx,
		[]gen.Condition{
			query.Q.Comment.ParentID.Eq(commentId),
		},
		[]field.Expr{
			query.Q.Comment.ID.Asc(),
		},
		limit,
		offset,
	)
}

func (r *PersistRepository) CountChildComments(ctx context.Context, commentId int64) (int64, error) {
	return query.Q.WithContext(ctx).Comment.Where(query.Comment.ParentID.Eq(commentId)).Count()
}

func (r *PersistRepository) GetById(ctx context.Context, commentId int64) (*model.Comment, error) {
	return dbtx.TxDoGetValue(ctx, func(tx *query.QueryTx) (*model.Comment, error) {
		return query.Q.WithContext(ctx).Comment.Where(query.Comment.ID.Eq(commentId)).First()
	})
}

func (r *PersistRepository) GetByIdList(ctx context.Context, commentIdList []int64) ([]*model.Comment, error) {
	return query.Q.WithContext(ctx).Comment.Where(query.Comment.ID.In(commentIdList...)).Find()
}

func (r *PersistRepository) count(ctx context.Context, conditions ...gen.Condition) (int64, error) {
	return query.Q.WithContext(ctx).Comment.Where(conditions...).Count()
}

func (r *PersistRepository) countWithGroup(ctx context.Context, targetFiled field.Int64, noChildComments bool, idList []int64) ([]*commentserviceiface.CountResult, error) {
	var result []*commentserviceiface.CountResult
	var conditions []gen.Condition
	conditions = append(conditions, targetFiled.In(idList...), query.Comment.IsDeleted.Is(false))
	if noChildComments {
		conditions = append(conditions, query.Comment.ParentID.Eq(0))
	}

	err := query.Q.WithContext(ctx).Comment.Select(
		targetFiled.As("id"),
		targetFiled.Count().As("count"),
	).Where(
		conditions...,
	).Group(
		targetFiled,
	).Scan(&result)
	return result, err
}

func (r *PersistRepository) CountByVideoId(ctx context.Context, videoId []int64, noChildComments bool) ([]*commentserviceiface.CountResult, error) {
	return r.countWithGroup(ctx, query.Q.Comment.VideoID, noChildComments, videoId)
}

func (r *PersistRepository) CountByUserId(ctx context.Context, userId []int64) ([]*commentserviceiface.CountResult, error) {
	return r.countWithGroup(ctx, query.Q.Comment.UserID, false, userId)
}

func (r *PersistRepository) CountParentCommentByVideoId(ctx context.Context, videoId int64) (int64, error) {
	return r.count(
		ctx,
		query.Comment.VideoID.Eq(videoId),
		query.Comment.ParentID.Eq(0),
		query.Comment.IsDeleted.Is(false),
	)
}

func (r *PersistRepository) CountByParentId(ctx context.Context, parentId int64) (int64, error) {
	return r.count(ctx, query.Comment.ParentID.Eq(parentId), query.Comment.IsDeleted.Is(false))
}

var _ repoiface.CommentRepository = (*PersistRepository)(nil)
