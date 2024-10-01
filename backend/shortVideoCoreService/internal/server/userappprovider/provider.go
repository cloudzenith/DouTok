package userappprovider

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/application/userapp"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/conf"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/data/userdata"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/service/userdomain"
)

func InitUserApplication(config *conf.Config) *userapp.UserApplication {
	userRepo := userdata.NewUserRepo()
	userUsecase := userdomain.NewUserUsecase(userRepo)
	userApp := userapp.NewUserApplication(userUsecase)
	return userApp
}
