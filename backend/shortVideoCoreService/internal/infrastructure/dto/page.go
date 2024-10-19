package dto

import v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"

type PaginationRequest struct {
	PageNum    int32        `json:"page_num"`
	PageSize   int32        `json:"page_size"`
	SortFields []*SortField `json:"sort_fields"`
}

type SortField struct {
	Field string `json:"field"`
	Order int32  `json:"order"` // 0: asc, 1: desc
}

type PaginationResponse struct {
	Page  int64 `json:"page"`
	Total int64 `json:"total"`
	Count int64 `json:"count"`
}

func FromPBSortField(s *v1.SortField) *SortField {
	return &SortField{
		Field: s.Field,
		Order: int32(s.Order),
	}
}

func FromPBPaginationRequest(p *v1.PaginationRequest) *PaginationRequest {
	sortFields := make([]*SortField, 0, len(p.Sort))
	for _, s := range p.Sort {
		sortFields = append(sortFields, FromPBSortField(s))
	}
	return &PaginationRequest{
		PageNum:    p.Page,
		PageSize:   p.Size,
		SortFields: sortFields,
	}
}

func ToPBPaginationResponse(p *PaginationResponse) *v1.PaginationResponse {
	return &v1.PaginationResponse{
		Page:  int32(p.Page),
		Total: int32(p.Total),
		Count: int32(p.Count),
	}
}
