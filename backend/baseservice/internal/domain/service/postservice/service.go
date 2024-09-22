package postservice

import (
	"context"
	"fmt"
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/entity/template"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/domain/repoiface"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"strings"
)

type PostService struct {
	templateRepo        repoiface.TemplateRepository
	thirdMsgSendService repoiface.ThirdMessageSendService
}

func New(templateRepo repoiface.TemplateRepository, thirdMsgSendService repoiface.ThirdMessageSendService) *PostService {
	return &PostService{
		templateRepo:        templateRepo,
		thirdMsgSendService: thirdMsgSendService,
	}
}

func (p *PostService) CreateTemplate(ctx context.Context, template *template.Template) (int64, error) {
	template.GenerateId()
	model := template.ToModel()
	err := p.templateRepo.Create(ctx, model)
	return model.ID, err
}

func (p *PostService) UpdateTemplate(ctx context.Context, template *template.Template) error {
	t, err := p.GetTemplateById(ctx, template.ID)
	if t == nil || err != nil {
		log.Context(ctx).Errorf("template id not exist: %v", err)
		return err
	}

	return p.templateRepo.ModifyById(ctx, template.ToModel())
}

func (p *PostService) GetTemplateById(ctx context.Context, id int64) (*template.Template, error) {
	t, err := p.templateRepo.GetById(ctx, id)
	if err != nil {
		log.Context(ctx).Errorf("get template failed: %v", err)
		return nil, err
	}

	return template.NewWithModel(t), nil
}

func (p *PostService) parse(fieldName string, operator api.SearchOperator, value string) (gen.Condition, error) {
	switch fieldName {
	case "title":
		return NewTemplateSearchFiled(query.Q.Template.Title, fieldName, operator).ToGormCondition(value)
	case "content":
		return NewTemplateSearchFiled(query.Q.Template.Content, fieldName, operator).ToGormCondition(value)
	default:
		return nil, fmt.Errorf("unsupported field name %s", fieldName)
	}
}

func (p *PostService) parseListTemplateSearchFields(searchFields ...*api.SearchField) ([]gen.Condition, error) {
	result := make([]gen.Condition, len(searchFields))
	for i := 0; i < len(searchFields); i++ {
		fieldName := searchFields[i].Field
		condition, err := p.parse(fieldName, searchFields[i].Operator, searchFields[i].Value)
		if err != nil {
			return nil, err
		}

		result[i] = condition
	}

	return result, nil
}

func (p *PostService) ListTemplate(ctx context.Context, pageRequest *api.PaginationRequest, searchFields ...*api.SearchField) ([]*template.Template, error) {
	page, size := pageRequest.GetPage(), pageRequest.GetSize()
	offset := (page - 1) * size
	limit := size

	conditions, err := p.parseListTemplateSearchFields(searchFields...)
	if err != nil {
		log.Context(ctx).Errorf("parse search fields failed: %v", err)
		return nil, err
	}

	result, err := p.templateRepo.Query(ctx, int(limit), int(offset), conditions...)
	if err != nil {
		log.Context(ctx).Errorf("query template failed: %v", err)
		return nil, err
	}

	templates := make([]*template.Template, len(result))
	for i := 0; i < len(result); i++ {
		templates[i] = template.NewWithModel(result[i])
	}
	return templates, nil
}

func (p *PostService) RemoveTemplate(ctx context.Context, id int64) error {
	if err := p.templateRepo.RemoveById(ctx, id); err != nil {
		log.Context(ctx).Errorf("remove template failed: %v", err)
		return err
	}

	return nil
}

func (p *PostService) concatMessageContent(content string, data map[string]string) string {
	for k, v := range data {
		content = strings.Replace(content, fmt.Sprintf("{%s}", k), v, -1)
	}

	return content
}

func (p *PostService) SendWithTemplate(ctx context.Context, sendType api.PostType, template *template.Template, to, title string, data map[string]string) error {
	content := p.concatMessageContent(template.Content, data)
	if err := p.Send(ctx, sendType, to, title, content); err != nil {
		log.Context(ctx).Errorf("send message failed: %v", err)
		return err
	}

	return nil
}

func (p *PostService) Send(ctx context.Context, sendType api.PostType, to, title, content string) error {
	if err := p.thirdMsgSendService.Send(ctx, sendType, to, title, content); err != nil {
		log.Context(ctx).Errorf("send message failed: %v", err)
		return err
	}

	return nil
}
