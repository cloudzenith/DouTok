package server

import (
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/middleware"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/collectionappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/commentappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/favoriteappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/followappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/userappprovider"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/server/videoappprovider"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

func NewGRPCServer(config *conf.Config) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			tracing.Server(),
			validate.Validator(),
			// 此处依赖的全局 logger 会跟随 launcher 的配置而变化
			logging.Server(log.GetLogger()),
			middleware.RequestMonitor(),
		),
	}

	if config.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(config.Server.Grpc.Addr))
	}

	srv := grpc.NewServer(opts...)

	v1.RegisterUserServiceServer(srv, userappprovider.InitUserApplication(config))
	v1.RegisterVideoServiceServer(srv, videoappprovider.InitVideoApplication(config))
	v1.RegisterCollectionServiceServer(srv, collectionappprovider.InitCollectionApplication())
	v1.RegisterCommentServiceServer(srv, commentappprovider.InitCommentApplication())
	v1.RegisterFavoriteServiceServer(srv, favoriteappprovider.InitFavoriteApp())
	v1.RegisterFollowServiceServer(srv, followappprovider.InitFollowApp())
	return srv
}
