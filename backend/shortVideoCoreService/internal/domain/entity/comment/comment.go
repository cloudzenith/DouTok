package comment

import (
	"encoding/json"
	"github.com/cloudzenith/DouTok/backend/gopkgs/snowflakeutil"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/utils"
	"time"
)

// FirstCommentsNumber 直接记录在Comment中的子评论个数
const FirstCommentsNumber = 3

type Comment struct {
	Id            int64
	VideoId       int64
	UserId        int64
	ParentId      *int64
	ToUserId      *int64
	Content       string
	Date          string
	CreateTime    time.Time
	Comments      []*Comment // 子评论
	ChildNumbers  int64      // 子评论个数
	FirstComments []*Comment // 最初的x条子评论
}

func New(options ...Option) *Comment {
	c := &Comment{}
	for _, option := range options {
		option(c)
	}

	return c
}

func NewWithModel(m *model.Comment) *Comment {
	var firstComments []*Comment
	_ = json.Unmarshal([]byte(m.FirstComments), &firstComments)

	return &Comment{
		Id:            m.ID,
		VideoId:       m.VideoID,
		UserId:        m.UserID,
		ParentId:      m.ParentID,
		ToUserId:      m.ToUserID,
		Content:       m.Content,
		Date:          utils.ParseToDateString(m.CreateTime),
		CreateTime:    m.CreateTime,
		FirstComments: firstComments,
	}
}

func (c *Comment) SetId() {
	c.Id = snowflakeutil.GetSnowflakeId()
}

func (c *Comment) ToModel() *model.Comment {
	v, _ := json.Marshal(c.FirstComments)

	return &model.Comment{
		ID:            c.Id,
		VideoID:       c.VideoId,
		UserID:        c.UserId,
		ParentID:      c.ParentId,
		ToUserID:      c.ToUserId,
		Content:       c.Content,
		FirstComments: string(v),
	}
}

func (c *Comment) GetChildCommentProto() []*v1.Comment {
	var result []*v1.Comment
	if len(c.FirstComments) == 0 {
		return result
	}

	for _, item := range c.FirstComments {
		result = append(result, item.ToProto())
	}

	return result
}

func (c *Comment) ToProto() *v1.Comment {
	return &v1.Comment{
		Id:          c.Id,
		VideoId:     c.VideoId,
		Content:     c.Content,
		Date:        utils.ParseToDateString(c.CreateTime),
		ReplyCount:  "0",
		UserId:      c.UserId,
		ParentId:    *c.ParentId,
		ReplyUserId: *c.ToUserId,
		Comments:    c.GetChildCommentProto(),
	}
}

func (c *Comment) AddFirstChildComments(child *Comment) {
	if len(c.FirstComments) >= FirstCommentsNumber {
		return
	}

	c.FirstComments = append(c.FirstComments, child)
}
