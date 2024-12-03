package request

import (
	"blog-service/constants"
	"blog-service/models"
)

// PaginationRequest represents the request parameters for pagination.
type PaginationRequest struct {
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

// NewPaginationRequest creates a new PaginationRequest with the specified values.
func NewPaginationRequest(page, pageSize int, sortBy, sortOrder string) *PaginationRequest {
	req := &PaginationRequest{
		Page:      page,
		PageSize:  pageSize,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
	req.SetDefaults() // Set defaults if necessary
	return req
}

// GetOffset calculates the offset for pagination based on the current page and page size.
func (p *PaginationRequest) GetOffset() int {
	return (p.Page - 1) * p.PageSize
}

// SetDefaults sets default values for pagination parameters if they are not provided.
func (p *PaginationRequest) SetDefaults() {
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

// Validate validates the pagination request parameters.
func (p *PaginationRequest) Validate() error {
	if p.Page < 1 {
		return models.ErrInvalidPage
	}
	if p.PageSize < 1 || p.PageSize > constants.MaxPageSize {
		return models.ErrInvalidPageSize
	}
	if p.SortOrder != constants.SortOrderAsc && p.SortOrder != constants.SortOrderDesc {
		return models.ErrInvalidSortOrder
	}
	return nil
}
