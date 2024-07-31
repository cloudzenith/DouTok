package do

import "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/model"

type User struct {
	ID              int64
	AccountID       int64
	Name            string
	Avatar          string
	BackgroundImage string
	Signature       string
	Mobile          string
	Email           string
	Password        string
}

func (u *User) ToUserModel() model.User {
	return model.User{
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

func FromUserModel(user model.User) *User {
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
