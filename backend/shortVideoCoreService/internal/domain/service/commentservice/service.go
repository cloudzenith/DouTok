package commentservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/commentserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/comment"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/repoiface"
	"github.com/go-kratos/kratos/v2/log"
)

type Service struct {
	comment repoiface.CommentRepository
}

func New() *Service {
	return &Service{}
}

func (s *Service) CreateComment(ctx context.Context, c *comment.Comment) (*comment.Comment, error) {
	err := s.comment.Create(ctx, c.ToModel())
	if err != nil {
		log.Context(ctx).Errorf("create comment failed, err: %v", err)
		return nil, err
	}

	return c, nil
}

func (s *Service) RemoveComment(ctx context.Context, userId, commentId int64) error {
	return s.comment.RemoveById(ctx, commentId)
}

func (s *Service) ListComment4Video(ctx context.Context, videoId int64, limit, offset int) ([]*comment.Comment, error) {
	comments, err := s.comment.ListByVideoId(ctx, videoId, limit, offset)
	if err != nil {
		log.Context(ctx).Errorf("list comment for video failed, err: %v", err)
		return nil, err
	}

	childCommentMap := make(map[int64][]*comment.Comment)
	for _, c := range comments {
		if c.ParentID != nil {
			if _, ok := childCommentMap[*c.ParentID]; !ok {
				childCommentMap[*c.ParentID] = make([]*comment.Comment, 0)
			}

			childCommentMap[*c.ParentID] = append(childCommentMap[*c.ParentID], comment.NewWithModel(c))
		}
	}

	var parentComments []*comment.Comment
	for _, c := range comments {
		if c.ParentID == nil {
			commentDO := comment.NewWithModel(c)
			commentDO.Comments = childCommentMap[c.ID]
			parentComments = append(parentComments, commentDO)
		}
	}

	return parentComments, nil
}

func (s *Service) GetCommentById(ctx context.Context, commentId int64) (*comment.Comment, error) {
	c, err := s.comment.GetById(ctx, commentId)
	if err != nil {
		log.Context(ctx).Errorf("get comment by id failed, err: %v", err)
		return nil, err
	}

	return comment.NewWithModel(c), nil
}

func (s *Service) CountComment4Video(ctx context.Context, videoId int64) (int64, error) {
	return s.comment.CountByVideoId(ctx, videoId)
}

func (s *Service) CountComment4User(ctx context.Context, userId int64) (int64, error) {
	return s.comment.CountByUserId(ctx, userId)
}

var _ commentserviceiface.CommentService = (*Service)(nil)
