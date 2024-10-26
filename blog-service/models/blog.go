package models

import "blog-service/constants"

type BlogListRequest struct {
	BaseRequest
	AuthorId *int    `json:"author_id,omitempty"`
	Category *string `json:"category,omitempty"`
}

func NewBlogListRequest() *BlogListRequest {
	return &BlogListRequest{
		BaseRequest: NewBaseRequest(),
	}
}

// SetDefaults sets default values for BlogListRequest
func (r *BlogListRequest) SetDefaults() {
	r.BaseRequest.SetDefaults()
	// Add any blog-specific default values here
}

// Validate validates the BlogListRequest
func (r *BlogListRequest) Validate() error {
	// first validate base fields
	if err := r.ValidateBase(); err != nil {
		return err
	}
	// validate blog specific fields
	if !isValidBlogSortField(r.SortBy) {
		r.SortBy = constants.DefaultSortBy
	}
	return nil

}

func (r *BlogListRequest) ValidateBase() error {

	if r.PageSize < 1 || r.PageSize > constants.MaxPageSize {
		return ErrInvalidPageSize
	}
	if r.SortOrder != constants.SortOrderAsc && r.SortOrder != constants.SortOrderDesc {
		return ErrInvalidSortOrder
	}
	return nil
}

func isValidBlogSortField(field string) bool {
	validFields := constants.ValidSortFields["blog"]
	for _, validField := range validFields {
		if field == validField {
			return true
		}
	}
	return false
}
