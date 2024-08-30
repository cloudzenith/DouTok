package baseservice

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/gopkgs/components/consulx"
)

type Handler struct {
	account api.AccountServiceClient
}

func NewHandler() (*Handler, error) {
	conn, err := consulx.GetGrpcConn(context.Background(), "discovery:///base-service")
	if err != nil {
		return nil, err
	}
	return &Handler{
		account: api.NewAccountServiceClient(conn),
	}, nil
}

func (h *Handler) Register(ctx context.Context, in *api.RegisterRequest) (*api.RegisterResponse, error) {
	return h.account.Register(ctx, in)
}

func (h *Handler) CheckAccount(ctx context.Context, in *api.CheckAccountRequest) (*api.CheckAccountResponse, error) {
	return h.account.CheckAccount(ctx, in)
}
