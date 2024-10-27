package services

import (
	"blog-service/models/schema"
	"blog-service/repositories"

	"github.com/sirupsen/logrus"
)

// BlogService ...
type BlogService interface {
	GetBlogs(pageNum, pageSize int) ([]schema.Blog, int, error)
	CreateBlog(blog *schema.Blog) error
	UpdateBlog(blog *schema.Blog) error
	DeleteBlog(blog *schema.Blog) error
	GetBlogById(blogId int) (*schema.Blog, error)
}

type blogService struct {
	repo repositories.BlogRepository
	log  *logrus.Logger
}

func (b blogService) GetBlogs(pageNum, pageSize int) ([]schema.Blog, int, error) {
	//TODO implement me
	panic("implement me")
}

func (b blogService) CreateBlog(blog *schema.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) UpdateBlog(blog *schema.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) DeleteBlog(blog *schema.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) GetBlogById(blogId int) (*schema.Blog, error) {
	//TODO implement me
	panic("implement me")
}

// NewBlogService creates  a new instance of BlogService
func NewBlogService(repo repositories.BlogRepository, l *logrus.Logger) BlogService {
	return &blogService{
		repo: repo,
		log:  l,
	}

}
