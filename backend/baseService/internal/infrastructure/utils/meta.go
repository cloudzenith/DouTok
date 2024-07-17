package utils

import "github.com/cloudzenith/DouTok/backend/baseService/api"

func GetSuccessMeta() *api.Metadata {
	return &api.Metadata{
		BizCode: 0,
		Message: "success",
	}
}

func GetMetaWithError(err error) *api.Metadata {
	return &api.Metadata{
		BizCode: -1,
		Message: err.Error(),
	}
}

func GetMetaWithErrorString(err string) *api.Metadata {
	return &api.Metadata{
		BizCode: -1,
		Message: err,
	}
}
