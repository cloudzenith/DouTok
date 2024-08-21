package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/etcdx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Adapter struct {
	user v1.UserServiceClient
}

func New() *Adapter {
	etcdClient := etcdx.GetClient(context.Background())
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///sv-core-service"),
		grpc.WithDiscovery(etcd.New(etcdClient)),
	)
	if err != nil {
		panic(err)
	}

	return &Adapter{
		user: v1.NewUserServiceClient(conn),
	}
}
