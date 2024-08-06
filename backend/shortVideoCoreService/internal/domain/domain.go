package domain

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewUserUsecase, NewBaseServiceClient)

func NewBaseServiceClient(thirdParty *conf.ThirdParty, logger log.Logger) (api.AccountServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(thirdParty.BaseService.Address),
	)
	if err != nil {
		return nil, err
	}
	return api.NewAccountServiceClient(conn), nil
}
