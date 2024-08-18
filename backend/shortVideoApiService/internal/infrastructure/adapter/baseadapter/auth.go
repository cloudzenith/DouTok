package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/internal/infrastructure/utils/respcheck"
)

func (a *Adapter) CreateVerificationCode(ctx context.Context, bits, expiredSeconds int64) (int64, error) {
	req := &api.CreateVerificationCodeRequest{
		Bits:       bits,
		ExpireTime: expiredSeconds * 1000,
	}

	resp, err := a.auth.CreateVerificationCode(ctx, req)
	return respcheck.CheckT[int64, *api.Metadata](
		resp, err,
		func() int64 {
			return resp.GetVerificationCodeId()
		},
	)
}

func (a *Adapter) ValidateVerificationCode(ctx context.Context, codeId int64, code string) error {
	req := &api.ValidateVerificationCodeRequest{
		VerificationCodeId: codeId,
		Code:               code,
	}

	resp, err := a.auth.ValidateVerificationCode(ctx, req)
	return respcheck.Check[*api.Metadata](resp, err)
}
