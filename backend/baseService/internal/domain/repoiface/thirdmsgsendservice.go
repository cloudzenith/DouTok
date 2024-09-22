package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
)

type ThirdMessageSendService interface {
	Send(ctx context.Context, sendType api.PostType, to, title, content string) error
}
