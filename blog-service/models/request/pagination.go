package request

import (
	"blog-service/constants"
	"blog-service/models"
)

type PaginationReq struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

// PaginationResponse represents common response metadata
type PaginationResponse struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
}

func NewPaginationReq() *PaginationReq {
	return &PaginationReq{
		Page:      constants.DefaultPage,
		PageSize:  constants.DefaultPageSize,
		SortBy:    constants.DefaultSortBy,
		SortOrder: constants.DefaultSortOrder,
	}
}

func (p *PaginationReq) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

func (p *PaginationReq) SetDefaults() {
	if p.PageSize < 1 || p.PageSize > constants.MaxPageSize {
		p.PageSize = constants.DefaultPageSize
	}
	if p.SortBy == "" {
		p.SortBy = constants.DefaultSortBy
	}
	if p.SortOrder == "" {
		p.SortOrder = constants.DefaultSortOrder
	}
}

func (r *BlogListReq) ValidateBase() error {
	if r.Page < 1 {
		return models.ErrInvalidPage
	}
	if r.PageSize < 1 || r.PageSize > constants.MaxPageSize {
		return models.ErrInvalidPageSize
	}

	if r.SortOrder != constants.SortOrderAsc && r.SortOrder != constants.SortOrderDesc {
		return models.ErrInvalidSortOrder
	}
	return nil
}
