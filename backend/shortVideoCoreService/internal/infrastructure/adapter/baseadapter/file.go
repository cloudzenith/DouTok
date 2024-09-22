package baseadapter

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
)

func (a *Adapter) PreSignGet4Forever(ctx context.Context) error {
	req := &api.PreSignGetRequest{}
	a.file.PreSignGet(ctx, req)
}
