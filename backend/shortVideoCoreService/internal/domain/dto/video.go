package dto

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/domain/entity"
	infra_dto "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/internal/infrastructure/dto"
)

type FeedShortVideoRequest struct {
	UserId     int64
	LatestTime int64
	FeedNum    int64
}

type FeedShortVideoResponse struct {
	Videos []*entity.Video
}

type PublishVideoRequest struct {
	UserId      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	VideoURL    string `json:"video_url"`
	CoverURL    string `json:"cover_url"`
}

type ListPublishedVideoRequest struct {
	UserId     int64
	LatestTime int64
	Pagination *infra_dto.PaginationRequest
}

type ListPublishedVideoResponse struct {
	Videos     []*entity.Video
	Pagination *infra_dto.PaginationResponse
}
