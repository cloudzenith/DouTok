package server

import (
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/userappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/videoappprovider"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(config *conf.Config, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			validate.Validator(),
		),
	}

	if config.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(config.Server.Grpc.Addr))
	}

	srv := grpc.NewServer(opts...)

	v1.RegisterUserServiceServer(srv, userappprovider.InitUserApplication(config, logger))
	v1.RegisterVideoServiceServer(srv, videoappprovider.InitVideoApplication(config, logger))
	return srv
}
