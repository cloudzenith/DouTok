// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameCollectionVideo = "collection_video"

// CollectionVideo mapped from table <collection_video>
type CollectionVideo struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey" json:"id"`
	CollectionID int64     `gorm:"column:collection_id;type:bigint;not null;index:collection_id_idx,priority:1" json:"collection_id"`
	UserID       int64     `gorm:"column:user_id;type:bigint;not null" json:"user_id"`
	VideoID      int64     `gorm:"column:video_id;type:bigint;not null" json:"video_id"`
	IsDeleted    bool      `gorm:"column:is_deleted;type:tinyint(1);not null;index:collection_id_idx,priority:2" json:"is_deleted"`
	CreateTime   time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime   time.Time `gorm:"column:update_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"update_time"`
}

// TableName CollectionVideo's table name
func (*CollectionVideo) TableName() string {
	return TableNameCollectionVideo
}
