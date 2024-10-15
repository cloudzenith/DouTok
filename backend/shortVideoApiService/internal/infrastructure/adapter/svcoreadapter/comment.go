package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/adapter/svcoreadapter/commentoptions"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func (a *Adapter) CreateComment(ctx context.Context, options ...commentoptions.CreateCommentOption) (*v1.Comment, error) {
	req := &v1.CreateCommentRequest{}
	for _, option := range options {
		option(req)
	}

	resp, err := a.comment.CreateComment(ctx, req)
	return respcheck.CheckT[*v1.Comment, *v1.Metadata](
		resp, err,
		func() *v1.Comment {
			return resp.Comment
		},
	)
}

func (a *Adapter) ListComment4Video(ctx context.Context, videoId int64, page, size int32) (*v1.ListComment4VideoResponse, error) {
	req := &v1.ListComment4VideoRequest{
		VideoId: videoId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.comment.ListComment4Video(ctx, req)
	return respcheck.CheckT[*v1.ListComment4VideoResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListComment4VideoResponse {
			return resp
		},
	)
}

func (a *Adapter) RemoveComment(ctx context.Context, commentId int64) error {
	req := &v1.RemoveCommentRequest{
		CommentId: commentId,
	}

	resp, err := a.comment.RemoveComment(ctx, req)
	return respcheck.Check[*v1.Metadata](resp, err)
}

func (a *Adapter) ListChildComments(ctx context.Context, commentId int64, page, size int32) (*v1.ListChildComment4CommentResponse, error) {
	req := &v1.ListChildComment4CommentRequest{
		CommentId: commentId,
		Pagination: &v1.PaginationRequest{
			Page: page,
			Size: size,
		},
	}

	resp, err := a.comment.ListChildComment4Comment(ctx, req)
	return respcheck.CheckT[*v1.ListChildComment4CommentResponse, *v1.Metadata](
		resp, err,
		func() *v1.ListChildComment4CommentResponse {
			return resp
		},
	)
}

func (a *Adapter) GetCommentById(ctx context.Context, commentId int64) (*v1.Comment, error) {
	req := &v1.GetCommentByIdRequest{
		CommentId: commentId,
	}

	resp, err := a.comment.GetCommentById(ctx, req)
	return respcheck.CheckT[*v1.Comment, *v1.Metadata](
		resp, err,
		func() *v1.Comment {
			return resp.Comment
		},
	)
}

func (a *Adapter) CountComments4Video(ctx context.Context, videoIdList []int64) (map[int64]int64, error) {
	req := &v1.CountComment4VideoRequest{
		VideoId: videoIdList,
	}

	resp, err := a.comment.CountComment4Video(ctx, req)
	return respcheck.CheckT[map[int64]int64, *v1.Metadata](
		resp, err,
		func() map[int64]int64 {
			result := make(map[int64]int64)
			for _, item := range resp.Results {
				result[item.Id] = item.Count
			}

			return result
		},
	)
}
