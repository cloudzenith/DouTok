package repoiface

import (
	"context"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/repositories/templaterepo"
	"gorm.io/gen"
)

type TemplateRepository interface {
	Create(ctx context.Context, template *models.Template) error
	ModifyById(ctx context.Context, template *models.Template) error
	Query(ctx context.Context, limit, offset int, conditions ...gen.Condition) ([]*models.Template, error)
	GetById(ctx context.Context, id int64) (*models.Template, error)
	RemoveById(ctx context.Context, id int64) error
}

var _ TemplateRepository = (*templaterepo.PersistRepository)(nil)
