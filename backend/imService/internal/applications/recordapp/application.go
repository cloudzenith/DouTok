package recordapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/imService/api/imapi"
)

type Application struct {
}

func New() *Application {
	return &Application{}
}

func (a *Application) Push(ctx context.Context, request *imapi.PushRequest) (*imapi.PushResponse, error) {
	return nil, nil
}

func (a *Application) Pull(ctx context.Context, request *imapi.PullRequest) (*imapi.PullResponse, error) {
	return nil, nil
}

func (a *Application) PullByAliveConnection(ctx context.Context, request *imapi.PullByAliveConnectionRequest) (*imapi.PullByAliveConnectionResponse, error) {
	return nil, nil
}

var _ imapi.RecordServiceServer = (*Application)(nil)
