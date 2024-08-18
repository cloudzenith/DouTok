package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/etcdx"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

type Adapter struct {
	account api.AccountServiceClient
	auth    api.AuthServiceClient
	file    api.FileServiceClient
}

func New() *Adapter {
	etcdClient := etcdx.GetClient(context.Background())
	conn, err := grpc.Dial(
		context.Background(),
		grpc.WithEndpoint("discovery:///provider"),
		grpc.WithDiscovery(etcd.New(etcdClient)),
	)
	if err != nil {
		panic(err)
	}

	return &Adapter{
		account: api.NewAccountServiceClient(conn),
		auth:    api.NewAuthServiceClient(conn),
		file:    api.NewFileServiceClient(conn),
	}
}
