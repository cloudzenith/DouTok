package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
)

type Adapter struct {
	file api.FileServiceClient
}

func New() *Adapter {
	conn, err := consulx.GetGrpcConn(context.Background(), "discovery:///base-service")
	if err != nil {
		panic(err)
	}

	return &Adapter{
		file: api.NewFileServiceClient(conn),
	}
}
