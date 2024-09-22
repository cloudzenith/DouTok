package templaterepo

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/query"
	"gorm.io/gen"
)

type PersistRepository struct {
}

func New() *PersistRepository {
	return &PersistRepository{}
}

func (p *PersistRepository) Create(ctx context.Context, template *models.Template) error {
	return query.Q.WithContext(ctx).Template.Create(template)
}

func (p *PersistRepository) ModifyById(ctx context.Context, template *models.Template) error {
	_, err := query.Q.WithContext(ctx).Template.Updates(template)
	return err
}

func (p *PersistRepository) list(ctx context.Context, queryFunc func(queryDo dal.QueryDo[*models.Template]) ([]*models.Template, error), conditions ...gen.Condition) ([]*models.Template, error) {
	conditions = append(conditions, query.Q.Template.IsDeleted.Is(false))
	handler := query.Q.WithContext(ctx).Template.Where(
		conditions...,
	)
	return queryFunc(handler)
}

func (p *PersistRepository) findByPage(offset, limit int) func(queryDo dal.QueryDo[*models.Template]) ([]*models.Template, error) {
	return func(queryDo dal.QueryDo[*models.Template]) ([]*models.Template, error) {
		result, _, err := queryDo.FindByPage(offset, limit)
		return result, err
	}
}

func (p *PersistRepository) find() func(queryDo dal.QueryDo[*models.Template]) ([]*models.Template, error) {
	return func(queryDo dal.QueryDo[*models.Template]) ([]*models.Template, error) {
		return queryDo.Find()
	}
}

func (p *PersistRepository) Query(ctx context.Context, limit, offset int, conditions ...gen.Condition) ([]*models.Template, error) {
	return p.list(
		ctx,
		p.findByPage(offset, limit),
		conditions...,
	)
}

func (p *PersistRepository) first(ctx context.Context, conditions ...gen.Condition) (*models.Template, error) {
	conditions = append(conditions, query.Q.Template.IsDeleted.Is(false))
	return query.Q.WithContext(ctx).Template.Where(
		conditions...,
	).First()
}

func (p *PersistRepository) GetById(ctx context.Context, id int64) (*models.Template, error) {
	return p.first(ctx, query.Q.Template.ID.Eq(id))
}

func (p *PersistRepository) removeById(ctx context.Context, templateId ...int64) error {
	if len(templateId) == 0 {
		return nil
	}

	templateList, err := p.list(
		ctx,
		p.find(),
		query.Q.Template.ID.In(templateId...),
	)

	if err != nil {
		return err
	}

	templateList4Removing := make([]*models.Template, 0, len(templateList))
	for _, template := range templateList {
		template.IsDeleted = true
		templateList4Removing = append(templateList4Removing, template)
	}

	_, err = query.Q.WithContext(ctx).Template.Updates(templateList4Removing)
	return err
}

func (p *PersistRepository) RemoveById(ctx context.Context, id int64) error {
	return p.removeById(ctx, id)
}
