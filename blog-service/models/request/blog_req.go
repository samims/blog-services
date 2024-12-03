package request

import "blog-service/constants"

// BlogListReq  is the request body for blog list
type BlogListReq struct {
	*PaginationRequest
	AuthorId *uint   `json:"author_id,omitempty"`
	Category *string `json:"category,omitempty"`
	Title    *string `json:"title,omitempty"`
}

// NewBlogListReq  returns a new instance of BlogListReq
func NewBlogListReq(page, pageSize int, sortBy, sortOrder string) *BlogListReq {
	return &BlogListReq{
		PaginationRequest: NewPaginationRequest(
			page,
			pageSize,
			sortBy,
			sortOrder,
		),
	}
}

// SetDefaults   sets default values for the request
func (r *BlogListReq) SetDefaults() {
	r.PaginationRequest.SetDefaults()
	// add any blog specific default values here
}

// Validate validates the request
func (r *BlogListReq) Validate() error {
	//  validate pagination
	if err := r.Validate(); err != nil {
		return err
	}
	//   validate blog specific fields
	if !isValidBlogSortField(r.SortBy) {
		// set  default sort field if  invalid
		// returning error is not  the best practice here
		r.SortBy = constants.DefaultSortBy
	}
	return nil
}

// isValidBlogSortField checks if the sortBy field is valid for blog list requests
func isValidBlogSortField(field string) bool {
	validFields := constants.ValidSortFields["blog"]

	for _, validField := range validFields {
		if field == validField {
			return true
		}
	}
	return false
}
