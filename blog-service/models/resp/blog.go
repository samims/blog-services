package resp

import (
	"fmt"
	"math"
)

// BlogPublicResp  is a response from the blog publications endpoint
type BlogPublicResp struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  uint   `json:"author"`
}

// BlogDetailResp  is a response from the blog detail endpoint
type BlogDetailResp struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Author    uint   `json:"author"`
	CreatedAt string `json:"create_at"`
	UpdatedAt string `json:"update_at"`
}

// BlogListPaginatedResp represents a paginated list of blog posts with items and pagination
type BlogListPaginatedResp struct {
	Items      []BlogPublicResp `json:"items"`
	Pagination PaginationResp   `json:"pagination"`
}

// PaginationResp represents pagination information
type PaginationResp struct {
	CurrentPage    int               `json:"current_page"`
	PageSize       int               `json:"page_size"`
	TotalItemCount int64             `json:"total_item_count"`
	TotalPages     int               `json:"total_pages"`
	HasMore        bool              `json:"has_more"`
	IsFirstPage    bool              `json:"is_first_page"`
	IsLastPage     bool              `json:"is_last_page"`
	Links          map[string]string `json:"links,omitempty"`
}

// BaseResponse represents a base response structure
type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// NewBaseResponse creates a new instance of BaseResponse
func NewBaseResponse() *BaseResponse {
	return &BaseResponse{
		Success: true,
		Data:    nil,
		Error:   "",
	}
}

// NewPaginationResp creates a new instance of PaginationResp
func NewPaginationResp(apiVersion string, currentPage, pageSize int, totalItemsCount int64) PaginationResp {
	totalPages := int(math.Ceil(float64(totalItemsCount) / float64(pageSize)))
	isFirstPage := currentPage == 1
	isLastPage := currentPage == totalPages

	links := map[string]string{}
	if !isFirstPage {
		links["prev"] = fmt.Sprintf("%s/%s?page=%d&size=%d", apiVersion, "blogs", currentPage-1, pageSize)
	} else {
		links["prev"] = ""
	}
	if !isLastPage {
		links["next"] = fmt.Sprintf("%s/%s?page=%d&size=%d", apiVersion, "blogs", currentPage+1, pageSize)
		links["last"] = fmt.Sprintf("%s/%s?page=%d&size=%d", apiVersion, "blogs", totalPages, pageSize)
		links["first"] = fmt.Sprintf("%s/%s?page=1&size=%d", apiVersion, "blogs", pageSize)
	}

	return PaginationResp{
		CurrentPage:    currentPage,
		PageSize:       pageSize,
		TotalItemCount: totalItemsCount,
		TotalPages:     totalPages,
		HasMore:        currentPage < totalPages,
		IsFirstPage:    isFirstPage,
		IsLastPage:     isLastPage,
		Links:          links,
	}
}
