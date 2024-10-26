package services

import (
	"blog-service/models"
	"blog-service/repositories"

	"github.com/sirupsen/logrus"
)

// BlogService ...
type BlogService interface {
	GetBlogs(pageNum, pageSize int) ([]models.Blog, int, error)
	CreateBlog(blog *models.Blog) error
	UpdateBlog(blog *models.Blog) error
	DeleteBlog(blog *models.Blog) error
	GetBlogById(blogId int) (*models.Blog, error)
}

type blogService struct {
	repo repositories.BlogRepository
	log  *logrus.Logger
}

func (b blogService) GetBlogs(pageNum, pageSize int) ([]models.Blog, int, error) {
	//TODO implement me
	panic("implement me")
}

func (b blogService) CreateBlog(blog *models.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) UpdateBlog(blog *models.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) DeleteBlog(blog *models.Blog) error {
	//TODO implement me
	panic("implement me")
}

func (b blogService) GetBlogById(blogId int) (*models.Blog, error) {
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
