package accountoptions

import "github.com/cloudzenith/DouTok/backend/baseService/api"

type RegisterOptions func(request *api.RegisterRequest)

func RegisterWithMobile(mobile string) RegisterOptions {
	return func(request *api.RegisterRequest) {
		request.Mobile = mobile
	}
}

func RegisterWithEmail(email string) RegisterOptions {
	return func(request *api.RegisterRequest) {
		request.Email = email
	}
}

func RegisterWithPassword(password string) RegisterOptions {
	return func(request *api.RegisterRequest) {
		request.Password = password
	}
}
