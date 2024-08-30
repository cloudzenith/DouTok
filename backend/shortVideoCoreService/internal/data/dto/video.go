package dto

import (
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/persistence/model"
)

type GetVideoListRequest struct {
	UserId            int64
	LatestTime        int64
	PaginationRequest *infra_dto.PaginationRequest
}

type GetVideoListResponse struct {
	Videos             []*model.Video
	PaginationResponse *infra_dto.PaginationResponse
}

type GetVideoFeedRequest struct {
	UserId     int64
	LatestTime int64
	Num        int64
}

type GetVideoFeedResponse struct {
	Videos []*model.Video
}
