package resp

type PaginationResp struct {
	CurrentPage    int   `json:"current_page"`
	PageSize       int   `json:"page_size"`
	TotalPage      int   `json:"total_page"`
	TotalItemCount int64 `json:"total_item_count"`
	HasMore        bool  `json:"has_more"`
	IsFirstPage    bool  `json:"is_first_page"`
	IsLastPage     bool  `json:"is_last_page"`
}
