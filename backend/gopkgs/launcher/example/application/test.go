package application

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/launcher/example/api"
)

type Application struct{}

func (a Application) Test(ctx context.Context, request *api.TestRequest) (*api.TestResponse, error) {
	return &api.TestResponse{
		Test: request.Test + request.Test,
	}, nil
}

var _ api.TestServiceServer = (*Application)(nil)
var _ api.TestServiceHTTPServer = (*Application)(nil)
