package thirdparty

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	kgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"google.golang.org/grpc"
)

type BaseService struct {
	AccountServiceClient *AccountServiceClient
}

func NewBaseService(config *conf.ThirdParty, logger log.Logger) (*BaseService, error) {
	conn, err := kgrpc.DialInsecure(
		context.Background(),
		kgrpc.WithEndpoint(config.BaseService.Address),
		kgrpc.WithMiddleware(
			logging.Client(logger),
		),
	)
	if err != nil {
		return nil, err
	}
	accountServiceClient := NewAccountServiceClient(conn, logger)
	return &BaseService{
		AccountServiceClient: accountServiceClient,
	}, nil
}

type AccountServiceClient struct {
	client api.AccountServiceClient
}

func (c *AccountServiceClient) Register(ctx context.Context, in *api.RegisterRequest) (*api.RegisterResponse, error) {
	return c.client.Register(ctx, in)
}

func (c *AccountServiceClient) CheckAccount(ctx context.Context, in *api.CheckAccountRequest) (*api.CheckAccountResponse, error) {
	return c.client.CheckAccount(ctx, in)
}

func NewAccountServiceClient(conn *grpc.ClientConn, logger log.Logger) *AccountServiceClient {
	client := api.NewAccountServiceClient(conn)
	return &AccountServiceClient{
		client: client,
	}
}
