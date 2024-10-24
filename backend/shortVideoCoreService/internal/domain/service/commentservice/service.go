package commentservice

import (
	"context"
	"github.com/TremblingV5/box/dbtx"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/commentserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/comment"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/go-kratos/kratos/v2/log"
	"time"
)

type Service struct {
	comment repoiface.CommentRepository
}

func New(comment repoiface.CommentRepository) *Service {
	return &Service{
		comment: comment,
	}
}

func (s *Service) CreateComment(ctx context.Context, c *comment.Comment) (cmt *comment.Comment, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	// 创建评论
	do := c.ToModel()
	err = s.comment.Create(ctx, do)
	if err != nil {
		log.Context(ctx).Errorf("create comment failed, err: %v", err)
		return nil, err
	}

	// 存在父评论
	if c.ParentId != nil && *c.ParentId > 0 {
		var parent *model.Comment
		parent, err = s.comment.GetById(ctx, *c.ParentId)
		if err != nil {
			log.Context(ctx).Errorf("get parent comment failed, err: %v", err)
			return nil, err
		}

		parentDO := comment.NewWithModel(parent)
		do.CreateTime = time.Now()
		parentDO.AddFirstChildComments(comment.NewWithModel(do))
		err = s.comment.Update(ctx, parentDO.ToModel())
		if err != nil {
			log.Context(ctx).Errorf("update parent comment failed, err: %v", err)
			return nil, err
		}
	}

	return c, nil
}

func (s *Service) RemoveComment(ctx context.Context, userId, commentId int64) error {
	return s.comment.RemoveById(ctx, commentId)
}

func (s *Service) ListComment4Video(ctx context.Context, videoId int64, limit, offset int) (*commentserviceiface.ListCommentsResult, error) {
	comments, err := s.comment.ListParentCommentByVideoId(ctx, videoId, limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("list comment for video failed, err: %v", err)
		return nil, err
	}

	countResult, err := s.comment.CountByVideoId(ctx, []int64{videoId}, true)
	if err != nil || len(countResult) != 1 {
		log.Context(ctx).Errorf("count parent comment by video id failed, err: %v", err)
		return nil, err
	}

	count := countResult[0].Count

	// FIXME: 此处写法不够合理，先实现功能，后续需要优化此处写法
	var parentComments []*comment.Comment
	for _, c := range comments {
		childCommentsNum, e := s.comment.CountChildComments(ctx, c.ID)
		if e != nil {
			log.Context(ctx).Errorf("count child comments failed, err: %v", e)
			continue
		}

		commentDO := comment.NewWithModel(c)
		commentDO.ChildNumbers = childCommentsNum
		parentComments = append(parentComments, commentDO)
	}

	return &commentserviceiface.ListCommentsResult{
		Data:  parentComments,
		Total: count,
	}, nil
}

func (s *Service) ListChildComment(ctx context.Context, commentId int64, limit, offset int) (*commentserviceiface.ListCommentsResult, error) {
	data, err := s.comment.ListChildCommentByCommentId(ctx, commentId, limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("failed to list child comments: %v", err)
		return nil, err
	}

	count, err := s.comment.CountByParentId(ctx, commentId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count child comments: %v", err)
		return nil, err
	}

	var result []*comment.Comment
	for _, item := range data {
		result = append(result, comment.NewWithModel(item))
	}

	return &commentserviceiface.ListCommentsResult{
		Data:  result,
		Total: count,
	}, nil
}

func (s *Service) GetCommentById(ctx context.Context, commentId int64) (cmt *comment.Comment, err error) {
	ctx, persist := dbtx.WithTXPersist(ctx)
	defer func() {
		persist(err)
	}()

	var c *model.Comment
	c, err = s.comment.GetById(ctx, commentId)
	if err != nil {
		log.Context(ctx).Errorf("get comment by id failed, err: %v", err)
		return nil, err
	}

	return comment.NewWithModel(c), nil
}

func (s *Service) CountComment4Video(ctx context.Context, videoId []int64) ([]*commentserviceiface.CountResult, error) {
	return s.comment.CountByVideoId(ctx, videoId, false)
}

func (s *Service) CountComment4User(ctx context.Context, userId []int64) ([]*commentserviceiface.CountResult, error) {
	return s.comment.CountByUserId(ctx, userId)
}

var _ commentserviceiface.CommentService = (*Service)(nil)
