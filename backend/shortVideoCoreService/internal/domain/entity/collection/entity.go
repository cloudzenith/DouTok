package collection

import (
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type Collection struct {
	ID          int64
	UserId      int64
	Title       string
	Description string
}

func (c *Collection) SetId() {
	c.ID = snowflakeutil.GetSnowflakeId()
}

func New(options ...Option) *Collection {
	c := &Collection{}
	for _, option := range options {
		option(c)
	}

	return c
}

func NewWithModel(m *model.Collection) *Collection {
	return &Collection{
		ID:          m.ID,
		UserId:      m.UserID,
		Title:       m.Title,
		Description: m.Description,
	}
}

func (c *Collection) ToModel() *model.Collection {
	return &model.Collection{
		ID:          c.ID,
		UserID:      c.UserId,
		Title:       c.Title,
		Description: c.Description,
	}
}
