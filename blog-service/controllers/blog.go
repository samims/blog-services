package controllers

import (
	"blog-service/services"
	"net/http"

	"github.com/sirupsen/logrus"
)

type BlogController interface {
	GetBlogList(w http.ResponseWriter, r *http.Request)
	GetBlog(w http.ResponseWriter, r *http.Request)
	CreateBlog(w http.ResponseWriter, r *http.Request)
	UpdateBlog(w http.ResponseWriter, r *http.Request)
	DeleteBlog(w http.ResponseWriter, r *http.Request)
}

type blogController struct {
	svc services.BlogService
	l   *logrus.Logger
}

func (b blogController) GetBlogList(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) ListBlogs(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) CreateBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) UpdateBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (b blogController) DeleteBlog(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewBlogController(svc services.BlogService, l *logrus.Logger) BlogController {
	return &blogController{
		svc: svc,
		l:   l,
	}
}
