package schema

import (
	"blog-service/models/resp"
	"time"
)

// Blog Post represents a blog post schema
type Blog struct {
	ID       uint      `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Author   uint      `json:"author"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

type BlogList []*Blog

// ToResponse converts a Blog entity to a BlogResponse.
// Used for preparing API responses.
func (b *Blog) ToResponse() *resp.BlogDetailResp {
	return &resp.BlogDetailResp{
		ID:       b.ID,
		Title:    b.Title,
		Content:  b.Content,
		Author:   b.Author,
		CreateAt: b.CreateAt.Format(time.RFC3339),
		UpdateAt: b.UpdateAt.Format(time.RFC3339),
	}
}

func (b *Blog) ToResponsePublic() *resp.BlogPublicResp {
	return &resp.BlogPublicResp{
		ID:      b.ID,
		Title:   b.Title,
		Content: b.Content,
		Author:  b.Author,
	}
}

// ToResponseList converts  a BlogList entity to a BlogResponseList.
// Used for preparing API responses.
func (bl BlogList) ToResponseList() []resp.BlogPublicResp {
	blogRespList := make([]resp.BlogPublicResp, 0, len(bl))

	for _, blog := range bl {
		blogRespList = append(blogRespList, *blog.ToResponsePublic())
	}
	return blogRespList
}

// ToPaginatedListResp converts blog list to a complete BlogListResponse with pagination
func (bl BlogList) ToPaginatedListResp(page, perPage int, total int64) *resp.BlogListPaginatedResp {
	paginatedRespList := &resp.BlogListPaginatedResp{
		Items:          bl.ToResponseList(),
		PaginationResp: resp.NewPaginationResp(page, perPage, total),
	}
	return paginatedRespList
}
