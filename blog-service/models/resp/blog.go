package resp

import (
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
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Author   uint   `json:"author"`
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"update_at"`
}

type BlogListResp struct {
	Items []BlogPublicResp `json:"items"`
}

type BlogListPaginatedResp struct {
	Items []BlogPublicResp `json:"items"`
	PaginationResp
}

type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{
		Success: true,
		Data:    nil,
		Error:   "",
	}
}

func NewPaginationResp(currentPage, pageSize int, totalItemsCount int64) PaginationResp {
	totalPages := int(math.Ceil(float64(totalItemsCount) / float64(pageSize)))
	isFirstPage := currentPage == 1
	isLastPage := currentPage == totalPages-1
	return PaginationResp{
		CurrentPage:    currentPage,
		PageSize:       pageSize,
		TotalItemCount: totalItemsCount,
		TotalPage:      totalPages,
		HasMore:        currentPage < totalPages,
		IsFirstPage:    isFirstPage,
		IsLastPage:     isLastPage,
	}
}
