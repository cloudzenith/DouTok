package svcoreadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

type Adapter struct {
	user v1.UserServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), "discovery:///sv-core-service")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		user: v1.NewUserServiceClient(conn),
	}
}
