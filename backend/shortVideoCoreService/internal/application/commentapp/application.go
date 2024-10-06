package commentapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/interface/commentserviceiface"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity/comment"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type Application struct {
	comment commentserviceiface.CommentService
	v1.UnimplementedCommentServiceServer
}

func New(comment commentserviceiface.CommentService) *Application {
	return &Application{
		comment: comment,
	}
}

func (a *Application) CreateComment(ctx context.Context, request *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	c := comment.New(
		comment.WithVideoId(request.VideoId),
		comment.WithUserId(request.UserId),
		comment.WithContent(request.Content),
		comment.WithParentId(request.ParentId),
		comment.WithToUserId(request.ReplyUserId),
	)
	c.SetId()

	result, err := a.comment.CreateComment(ctx, c)
	if err != nil {
		log.Context(ctx).Errorf("failed to create comment: %v", err)
		return &v1.CreateCommentResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CreateCommentResponse{
		Meta:    utils.GetSuccessMeta(),
		Comment: result.ToProto(),
	}, nil
}

func (a *Application) RemoveComment(ctx context.Context, request *v1.RemoveCommentRequest) (*v1.RemoveCommentResponse, error) {
	err := a.comment.RemoveComment(ctx, request.UserId, request.CommentId)
	if err != nil {
		log.Context(ctx).Errorf("failed to remove comment: %v", err)
		return &v1.RemoveCommentResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.RemoveCommentResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (a *Application) ListComment4Video(ctx context.Context, request *v1.ListComment4VideoRequest) (*v1.ListComment4VideoResponse, error) {
	limit, offset := utils.GetLimitOffset(
		int(request.Pagination.Size),
		(int(request.Pagination.Page)-1)*int(request.Pagination.Size),
	)

	comments, err := a.comment.ListComment4Video(
		ctx,
		request.VideoId,
		limit, offset,
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to list comments: %v", err)
		return &v1.ListComment4VideoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var commentsProto []*v1.Comment
	for _, c := range comments {
		commentsProto = append(commentsProto, c.ToProto())
	}

	return &v1.ListComment4VideoResponse{
		Meta:     utils.GetSuccessMeta(),
		Comments: commentsProto,
	}, nil
}

func (a *Application) ListChildComment4Comment(ctx context.Context, request *v1.ListChildComment4CommentRequest) (*v1.ListChildComment4CommentResponse, error) {
	limit, offset := utils.GetLimitOffset(
		int(request.Pagination.Size),
		(int(request.Pagination.Page)-1)*int(request.Pagination.Size),
	)

	comments, err := a.comment.ListChildComment(
		ctx,
		request.CommentId,
		limit, offset,
	)
	if err != nil {
		log.Context(ctx).Errorf("failed to list child comments: %v", err)
		return &v1.ListChildComment4CommentResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var commentsProto []*v1.Comment
	for _, c := range comments {
		commentsProto = append(commentsProto, c.ToProto())
	}

	return &v1.ListChildComment4CommentResponse{
		Meta:     utils.GetSuccessMeta(),
		Comments: commentsProto,
	}, nil
}

func (a *Application) GetCommentById(ctx context.Context, request *v1.GetCommentByIdRequest) (*v1.GetCommentByIdResponse, error) {
	cmt, err := a.comment.GetCommentById(ctx, request.CommentId)
	if err != nil {
		log.Context(ctx).Errorf("failed to get comment: %v", err)
		return &v1.GetCommentByIdResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.GetCommentByIdResponse{
		Meta:    utils.GetSuccessMeta(),
		Comment: cmt.ToProto(),
	}, nil
}

func (a *Application) CountComment4Video(ctx context.Context, request *v1.CountComment4VideoRequest) (*v1.CountComment4VideoResponse, error) {
	num, err := a.comment.CountComment4Video(ctx, request.VideoId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count comments: %v", err)
		return &v1.CountComment4VideoResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CountComment4VideoResponse{
		Meta:  utils.GetSuccessMeta(),
		Count: num,
	}, nil
}

func (a *Application) CountComment4User(ctx context.Context, request *v1.CountComment4UserRequest) (*v1.CountComment4UserResponse, error) {
	num, err := a.comment.CountComment4User(ctx, request.UserId)
	if err != nil {
		log.Context(ctx).Errorf("failed to count comments: %v", err)
		return &v1.CountComment4UserResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &v1.CountComment4UserResponse{
		Meta:  utils.GetSuccessMeta(),
		Count: num,
	}, nil
}

//var _ v1.CommentServiceServer = (*Application)(nil)
