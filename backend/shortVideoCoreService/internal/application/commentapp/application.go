package commentapp

import (
	"context"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) CreateComment(ctx context.Context, request *v1.CreateCommentRequest) (*v1.CreateCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) RemoveComment(ctx context.Context, request *v1.RemoveCommentRequest) (*v1.RemoveCommentResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) ListComment4Video(ctx context.Context, request *v1.ListComment4VideoRequest) (*v1.ListComment4VideoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) GetCommentById(ctx context.Context, request *v1.GetCommentByIdRequest) (*v1.GetCommentByIdResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) CountComment4Video(ctx context.Context, request *v1.CountComment4VideoRequest) (*v1.CountComment4VideoResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (a *Application) CountComment4User(ctx context.Context, request *v1.CountComment4UserRequest) (*v1.CountComment4UserResponse, error) {
	//TODO implement me
	panic("implement me")
}
