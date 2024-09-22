package template

import (
	"github.com/cloudzenith/DouTok/backend/baseService/api"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/dal/models"
	"github.com/cloudzenith/DouTok/backend/baseService/internal/infrastructure/utils"
	"time"
)

type Template struct {
	ID         int64
	Title      string
	Content    string
	IsDelete   bool
	CreateTime time.Time
	UpdateTime time.Time
}

func New(options ...Option) *Template {
	template := &Template{}
	for _, option := range options {
		option(template)
	}

	return template
}

func NewWithModel(model *models.Template) *Template {
	return &Template{
		ID:         model.ID,
		Title:      model.Title,
		Content:    model.Content,
		IsDelete:   model.IsDeleted,
		CreateTime: model.CreateTime,
		UpdateTime: model.UpdateTime,
	}
}

func (t *Template) ToDTO() *api.Template {
	return &api.Template{
		Id:         t.ID,
		Title:      t.Title,
		Content:    t.Content,
		CreateTime: t.CreateTime.UnixMilli(),
		UpdateTime: t.UpdateTime.UnixMilli(),
	}
}

func (t *Template) ToModel() *models.Template {
	return &models.Template{
		ID:         t.ID,
		Title:      t.Title,
		Content:    t.Content,
		IsDeleted:  t.IsDelete,
		CreateTime: t.CreateTime,
		UpdateTime: t.UpdateTime,
	}
}

func (t *Template) GenerateId() {
	t.ID = utils.GetSnowflakeId()
}

func (t *Template) Update(options ...Option) {
	for _, option := range options {
		option(t)
	}
}
