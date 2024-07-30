package thirdmsgadapter

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/go-kratos/kratos/v2/log"
)

type ThirdMsgAdapter struct {
}

func New() *ThirdMsgAdapter {
	return &ThirdMsgAdapter{}
}

func (t *ThirdMsgAdapter) Send(ctx context.Context, sendType api.PostType, to, title, content string) error {
	switch sendType {
	case api.PostType_SMS:
		return t.mockSendSms(ctx, to, title, content)
	case api.PostType_EMAIL:
		return t.mockSendEmail(ctx, to, title, content)
	default:
		return fmt.Errorf("unknown post type %s", sendType)
	}
}

func (t *ThirdMsgAdapter) mockSendSms(ctx context.Context, to, title, content string) error {
	log.Context(ctx).Infow(
		"msg", "mock send sms",
		"to", to,
		"title", title,
		"content", content,
	)
	return nil
}

func (t *ThirdMsgAdapter) mockSendEmail(ctx context.Context, to, title, content string) error {
	log.Context(ctx).Infow(
		"msg", "mock send email",
		"to", to,
		"title", title,
		"content", content,
	)
	return nil
}
