package accountoptions

import "github.com/cloudzenith/DouTok/backend/baseService/api"

type CheckAccountOption func(request *api.CheckAccountRequest)

func CheckAccountWithMobile(mobile string) CheckAccountOption {
	return func(request *api.CheckAccountRequest) {
		request.Mobile = mobile
	}
}

func CheckAccountWithEmail(email string) CheckAccountOption {
	return func(request *api.CheckAccountRequest) {
		request.Email = email
	}
}

func CheckAccountWithAccountId(accountId int64) CheckAccountOption {
	return func(request *api.CheckAccountRequest) {
		request.AccountId = accountId
	}
}

func CheckAccountWithPassword(password string) CheckAccountOption {
	return func(request *api.CheckAccountRequest) {
		request.Password = password
	}
}
