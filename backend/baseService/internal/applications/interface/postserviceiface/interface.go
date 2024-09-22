package postserviceiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/template"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/service/postservice"
)

type PostService interface {
	CreateTemplate(ctx context.Context, template *template.Template) (int64, error)
	UpdateTemplate(ctx context.Context, template *template.Template) error
	GetTemplateById(ctx context.Context, id int64) (*template.Template, error)
	ListTemplate(ctx context.Context, pageRequest *api.PaginationRequest, searchFields ...*api.SearchField) ([]*template.Template, error)
	RemoveTemplate(ctx context.Context, id int64) error
	SendWithTemplate(ctx context.Context, sendType api.PostType, template *template.Template, to, title string, data map[string]string) error
	Send(ctx context.Context, sendType api.PostType, to, title, content string) error
}

var _ PostService = (*postservice.PostService)(nil)
