package model

import "time"

const TableNameVideo = "video"

type Video struct {
	ID           int64     `gorm:"column:id;type:bigint(20);primaryKey"`
	UserID       int64     `gorm:"column:user_id;type:bigint(20);index"`
	Title        string    `gorm:"column:title;type:varchar(255)"`
	Description  string    `gorm:"column:description;type:varchar(255)"`
	VideoURL     string    `gorm:"column:video_url;type:varchar(255)"`
	CoverURL     string    `gorm:"column:cover_url;type:varchar(255)"`
	LikeCount    int64     `gorm:"column:like_count;type:bigint(20)"`
	CommentCount int64     `gorm:"column:comment_count;type:bigint(20)"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (*Video) TableName() string {
	return TableNameVideo
}
