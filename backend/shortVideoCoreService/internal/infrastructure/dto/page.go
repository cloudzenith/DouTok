package dto

import v1 "github.com/cloudzenith/DouTok/backend/shortVideoCoreService/api/v1"

type PaginationRequest struct {
	PageNum    int32        `json:"page_num"`
	PageSize   int32        `json:"page_size"`
	SortFields []*SortField `json:"sort_fields"`
}

func (p *PaginationRequest) getOffset() int32 {
	return (p.PageNum - 1) * p.PageSize
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
