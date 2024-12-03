package schema

import (
	"blog-service/models/resp"

	"time"
)

// Blog represents a blog post schema
type Blog struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	AuthorID  uint      `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BlogList represents a list of blog posts schema
type BlogList []Blog

// ToResponse converts a Blog entity to a BlogDetailResp.
// Used for preparing detailed API responses.
func (b *Blog) ToResponse() *resp.BlogDetailResp {
	return &resp.BlogDetailResp{
		ID:        b.ID,
		Title:     b.Title,
		Content:   b.Content,
		Author:    b.AuthorID,
		CreatedAt: b.CreatedAt.Format(time.RFC3339),
		UpdatedAt: b.UpdatedAt.Format(time.RFC3339),
	}
}

// ToResponsePublic converts a Blog entity to a BlogPublicResp.
// Used for preparing public API responses.
func (b *Blog) ToResponsePublic() resp.BlogPublicResp {
	return resp.BlogPublicResp{
		ID:      b.ID,
		Title:   b.Title,
		Content: b.Content,
		Author:  b.AuthorID,
	}
}

// ToResponseList converts a BlogList to a slice of BlogPublicResp.
// Used for preparing a list of public API responses.
func (bl BlogList) ToResponseList() []resp.BlogPublicResp {
	blogRespList := make([]resp.BlogPublicResp, len(bl)) // Preallocate slice

	for i, blog := range bl {
		blogRespList[i] = blog.ToResponsePublic()
	}
	return blogRespList
}

// ToPaginatedListResp converts a BlogList to a complete BlogListPaginatedResp with pagination.
// This method prepares a paginated response for API requests.
func (bl BlogList) ToPaginatedListResp(apiVersion string, page, perPage int, total int64) *resp.BlogListPaginatedResp {
	return &resp.BlogListPaginatedResp{
		Items:      bl.ToResponseList(),
		Pagination: resp.NewPaginationResp(apiVersion, page, perPage, total),
	}
}
