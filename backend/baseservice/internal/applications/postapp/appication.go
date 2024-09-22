package postapp

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/applications/interface/postserviceiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/template"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"github.com/go-kratos/kratos/v2/log"
)

type PostApplication struct {
	postService postserviceiface.PostService
}

func New(postService postserviceiface.PostService) *PostApplication {
	return &PostApplication{
		postService: postService,
	}
}

func (p *PostApplication) CreateTemplate(ctx context.Context, request *api.CreateTemplateRequest) (*api.CreateTemplateResponse, error) {
	template := template.New(
		template.WithTitle(request.Title),
		template.WithContent(request.Content),
	)
	_, err := p.postService.CreateTemplate(ctx, template)
	if err != nil {
		log.Context(ctx).Errorf("create template failed: %v", err)
		return &api.CreateTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.CreateTemplateResponse{
		Meta:     utils.GetSuccessMeta(),
		Template: template.ToDTO(),
	}, nil
}

func (p *PostApplication) UpdateTemplate(ctx context.Context, request *api.UpdateTemplateRequest) (*api.UpdateTemplateResponse, error) {
	t, err := p.postService.GetTemplateById(ctx, request.TemplateId)
	if err != nil {
		log.Context(ctx).Errorf("get template failed: %v", err)
		return &api.UpdateTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	t.Update(
		template.WithTitle(request.Title),
		template.WithContent(request.Content),
	)
	err = p.postService.UpdateTemplate(ctx, t)
	if err != nil {
		log.Context(ctx).Errorf("update template failed: %v", err)
		return &api.UpdateTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.UpdateTemplateResponse{
		Meta:     utils.GetSuccessMeta(),
		Template: t.ToDTO(),
	}, nil
}

func (p *PostApplication) ListTemplate(ctx context.Context, request *api.ListTemplateRequest) (*api.ListTemplateResponse, error) {
	templates, err := p.postService.ListTemplate(ctx, request.GetPaginationRequest(), request.GetSearchFields()...)
	if err != nil {
		log.Context(ctx).Errorf("list template failed: %v", err)
		return &api.ListTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	var dtos []*api.Template
	for _, t := range templates {
		dtos = append(dtos, t.ToDTO())
	}

	return &api.ListTemplateResponse{
		Meta:      utils.GetSuccessMeta(),
		Templates: dtos,
	}, nil
}

func (p *PostApplication) GetTemplate(ctx context.Context, request *api.GetTemplateRequest) (*api.GetTemplateResponse, error) {
	t, err := p.postService.GetTemplateById(ctx, request.TemplateId)
	if err != nil {
		log.Context(ctx).Errorf("get template failed: %v", err)
		return &api.GetTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.GetTemplateResponse{
		Meta:     utils.GetSuccessMeta(),
		Template: t.ToDTO(),
	}, nil
}

func (p *PostApplication) RemoveTemplate(ctx context.Context, request *api.RemoveTemplateRequest) (*api.RemoveTemplateResponse, error) {
	err := p.postService.RemoveTemplate(ctx, request.TemplateId)
	if err != nil {
		log.Context(ctx).Errorf("remove template failed: %v", err)
		return &api.RemoveTemplateResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.RemoveTemplateResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (p *PostApplication) SendSms(ctx context.Context, request *api.SendSmsRequest) (*api.SendSmsResponse, error) {
	t, err := p.postService.GetTemplateById(ctx, request.TemplateId)
	if err != nil {
		log.Context(ctx).Errorf("get template failed: %v", err)
		return &api.SendSmsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	err = p.postService.SendWithTemplate(ctx, api.PostType_SMS, t, request.To, "", request.Data)
	if err != nil {
		log.Context(ctx).Errorf("send sms failed: %v", err)
		return &api.SendSmsResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.SendSmsResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (p *PostApplication) SendEmail(ctx context.Context, request *api.SendEmailRequest) (*api.SendEmailResponse, error) {
	t, err := p.postService.GetTemplateById(ctx, request.TemplateId)
	if err != nil {
		log.Context(ctx).Errorf("get template failed: %v", err)
		return &api.SendEmailResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	err = p.postService.SendWithTemplate(ctx, api.PostType_EMAIL, t, request.To, request.EmailTitle, request.Data)
	if err != nil {
		log.Context(ctx).Errorf("send email failed: %v", err)
		return &api.SendEmailResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.SendEmailResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

func (p *PostApplication) Send(ctx context.Context, request *api.SendRequest) (*api.SendResponse, error) {
	err := p.postService.Send(ctx, request.PostType, request.To, request.Title, request.Content)
	if err != nil {
		log.Context(ctx).Errorf("send failed: %v", err)
		return &api.SendResponse{
			Meta: utils.GetMetaWithError(err),
		}, nil
	}

	return &api.SendResponse{
		Meta: utils.GetSuccessMeta(),
	}, nil
}

var _ api.PostServiceServer = (*PostApplication)(nil)
