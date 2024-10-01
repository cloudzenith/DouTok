package comment

import (
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"time"
)

type Comment struct {
	Id         int64
	VideoId    int64
	UserId     int64
	ParentId   *int64
	ToUserId   *int64
	Content    string
	Date       string
	CreateTime time.Time
	Comments   []*Comment // 子评论
}

func New(options ...Option) *Comment {
	c := &Comment{}
	for _, option := range options {
		option(c)
	}

	return c
}

func NewWithModel(m *model.Comment) *Comment {
	return &Comment{
		Id:         m.ID,
		VideoId:    m.VideoID,
		UserId:     m.UserID,
		ParentId:   m.ParentID,
		ToUserId:   m.ToUserID,
		Content:    m.Content,
		Date:       utils.ParseToDateString(m.CreateTime),
		CreateTime: m.CreateTime,
	}
}

func (c *Comment) SetId() {
	c.Id = snowflakeutil.GetSnowflakeId()
}

func (c *Comment) ToModel() *model.Comment {
	return &model.Comment{
		ID:       c.Id,
		VideoID:  c.VideoId,
		UserID:   c.UserId,
		ParentID: c.ParentId,
		ToUserID: c.ToUserId,
		Content:  c.Content,
	}
}
