package respcheck

import (
	"github.com/cloudzenith/DouTok/backend/shortVideoApiService/api/svapi"
	v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"
)

func ParseSvCorePagination(pagination *v1.PaginationResponse) *svapi.PaginationResponse {
	return &svapi.PaginationResponse{
		Page:  pagination.Page,
		Total: pagination.Total,
		Count: pagination.Count,
	}
}
