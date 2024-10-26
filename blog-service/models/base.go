package models

import (
	"blog-service/constants"
)

// BaseRequest represents common request fields
type BaseRequest struct {
	Page      int    `json:"page" validate:"min=1"`
	PageSize  int    `json:"pageSize" validate:"min=1"`
	SortBy    string `json:"sortBy"`
	SortOrder string `json:"sortOrder"`
}

// NewBaseRequest creates a new BaseRequest with default values
func NewBaseRequest() BaseRequest {
	return BaseRequest{
		Page:      constants.DefaultPage,
		PageSize:  constants.DefaultPageSize,
		SortBy:    constants.DefaultSortBy,
		SortOrder: constants.DefaultSortOrder,
	}
}

// SetDefaults are default values for empty or invalid filed values
// TODO: ...
// SetDefaults sets default values if fields are empty or invalid
func (r *BaseRequest) SetDefaults() {
	if r.Page < 1 {
		r.Page = constants.DefaultPage
	}
	if r.PageSize < 1 || r.PageSize > constants.MaxPageSize {
		r.PageSize = constants.DefaultPageSize
	}
	if r.SortOrder != constants.SortOrderAsc && r.SortOrder != constants.SortOrderDesc {
		r.SortOrder = constants.DefaultSortOrder
	}
	if r.SortBy == "" {
		r.SortBy = constants.DefaultSortBy
	}
}

// BaseResponse represents common response fields
type BaseResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// NewBaseResponse creates a new BaseResponse with success status
func NewBaseResponse() BaseResponse {
	return BaseResponse{
		Success: true,
	}
}

type PageMetaData struct {
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	TotalCount  int64 `json:"total_count"`
	HasNext     bool  `json:"has_next"`
}

type PagedResponse struct {
	BaseResponse
	Metadata PageMetaData `json:"metadata"`
}
