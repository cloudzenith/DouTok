package commentserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/comment"
)

type ListCommentsResult struct {
	Data  []*comment.Comment
	Total int64
}

type CommentService interface {
	CreateComment(ctx context.Context, c *comment.Comment) (*comment.Comment, error)
	RemoveComment(ctx context.Context, userId, commentId int64) error
	ListComment4Video(ctx context.Context, videoId int64, limit, offset int) (*ListCommentsResult, error)
	ListChildComment(ctx context.Context, commentId int64, limit, offset int) (*ListCommentsResult, error)
	GetCommentById(ctx context.Context, commentId int64) (*comment.Comment, error)
	CountComment4Video(ctx context.Context, videoId int64) (int64, error)
	CountComment4User(ctx context.Context, userId int64) (int64, error)
}
