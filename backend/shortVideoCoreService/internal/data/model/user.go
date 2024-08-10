package model

import (
	"time"
)

const TableNameUser = "user"

type User struct {
	ID              int64     `gorm:"column:id;type:bigint(20);primaryKey"`
	AccountID       int64     `gorm:"column:account_id;type:bigint(20);unique_index"`
	Mobile          string    `gorm:"column:mobile;type:varchar(20);unique_index:user_mobile_uidx"`
	Email           string    `gorm:"column:email;type:varchar(50);unique_index:user_email_uidx"`
	Name            string    `gorm:"column:name;type:varchar(50)"`
	Avatar          string    `gorm:"column:avatar;type:varchar(255)"`
	BackgroundImage string    `gorm:"column:background_image;type:varchar(255)"`
	Signature       string    `gorm:"column:signature;type:varchar(255)"`
	CreatedAt       time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt       time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (*User) TableName() string {
	return TableNameUser
}
