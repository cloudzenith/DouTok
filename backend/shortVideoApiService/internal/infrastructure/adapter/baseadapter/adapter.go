package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
)

type Adapter struct {
	account api.AccountServiceClient
	auth    api.AuthServiceClient
	file    api.FileServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), "discovery:///base-service")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		account: api.NewAccountServiceClient(conn),
		auth:    api.NewAuthServiceClient(conn),
		file:    api.NewFileServiceClient(conn),
	}
}
