package useroptions

import v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"

type UpdateUserInfoOption func(request *v1.UpdateUserInfoRequest)

func UpdateUserInfoWithUserId(userId int64) UpdateUserInfoOption {
	return func(request *v1.UpdateUserInfoRequest) {
		request.UserId = userId
	}
}

func UpdateUserInfoWithName(name string) UpdateUserInfoOption {
	return func(request *v1.UpdateUserInfoRequest) {
		request.Name = name
	}
}

func UpdateUserInfoWithAvatar(avatar string) UpdateUserInfoOption {
	return func(request *v1.UpdateUserInfoRequest) {
		request.Avatar = avatar
	}
}

func UpdateUserInfoWithBackgroundImage(backgroundImage string) UpdateUserInfoOption {
	return func(request *v1.UpdateUserInfoRequest) {
		request.BackgroundImage = backgroundImage
	}
}

func UpdateUserInfoWithSignature(signature string) UpdateUserInfoOption {
	return func(request *v1.UpdateUserInfoRequest) {
		request.Signature = signature
	}
}
