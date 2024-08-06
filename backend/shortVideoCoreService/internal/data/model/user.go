package model

import (
	"time"
)

type User struct {
	ID              int64     `gorm:"primarykey"`
	AccountID       int64     `gorm:"unique_index"`
	Mobile          string    `gorm:"size:20;unique_index"`
	Email           string    `gorm:"size:50;unique_index"`
	Name            string    `gorm:"size:50"`
	Avatar          string    `gorm:"size:255"`
	BackgroundImage string    `gorm:"size:255"`
	Signature       string    `gorm:"size:255"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
