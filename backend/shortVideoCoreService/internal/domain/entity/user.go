package entity

import (
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type User struct {
	ID              int64
	AccountID       int64
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	Mobile          string
	Email           string
}

func (u *User) ToUserResp() *v1.User {
	return &v1.User{
		Id:              u.ID,
		Name:            u.Name,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
		Mobile:          u.Mobile,
		Email:           u.Email,
	}
}

func (u *User) ToUserModel() *model.User {
	return &model.User{
		ID:              u.ID,
		AccountID:       u.AccountID,
		Mobile:          u.Mobile,
		Email:           u.Email,
		Name:            u.Name,
		Avatar:          u.Avatar,
		BackgroundImage: u.BackgroundImage,
		Signature:       u.Signature,
	}
}

func FromUserModel(user *model.User) *User {
	return &User{
		ID:              user.ID,
		AccountID:       user.AccountID,
		Mobile:          user.Mobile,
		Email:           user.Email,
		Name:            user.Name,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
	}
}
