package services

import (
	"blog-service/constants"
	"blog-service/logger"
	"blog-service/models/request"
	"blog-service/models/resp"
	"blog-service/models/schema"
	"blog-service/repositories"
	"context"
	"errors"
	"fmt"
)

// BlogService defines the methods for interacting with blog data.
type BlogService interface {
	GetAllBlogs(ctx context.Context, pageReq request.PaginationRequest) (*resp.BlogListPaginatedResp, error)
	CreateBlog(ctx context.Context, blog *schema.Blog) error
	UpdateBlog(ctx context.Context, blog *schema.Blog) error
	DeleteBlog(ctx context.Context, id int64, authUserID uint) error
	GetBlogById(ctx context.Context, blogId int64) (*schema.Blog, error)
	GetBlogsByAuthorID(ctx context.Context, authorID int64, pageReq request.PaginationRequest) (*resp.BlogListPaginatedResp, error)
}

// blogService is a concrete implementation of BlogService.
type blogService struct {
	blogRepo repositories.BlogRepository
	log      *logger.AppLogger
}

// NewBlogService creates a new instance of BlogService.
func NewBlogService(repo repositories.BlogRepository, logger *logger.AppLogger) BlogService {
	return &blogService{
		blogRepo: repo,
		log:      logger,
	}
}

// GetAllBlogs retrieves all blogs with pagination.
func (s *blogService) GetAllBlogs(ctx context.Context, pageReq request.PaginationRequest) (*resp.BlogListPaginatedResp, error) {
	s.log.WithContext(ctx).Infof("Fetching all blogs with pagination: %+v", pageReq)

	// Call the repository to get the blogs and total count
	blogs, totalCount, err := s.blogRepo.GetAllBlogs(ctx, pageReq)
	if err != nil {
		s.log.WithError(err).Error("Failed to fetch blogs from repository")
		return nil, fmt.Errorf("could not retrieve blogs: %w", err)
	}

	// Handle case where there are no blogs
	if totalCount == 0 {
		s.log.Info(ctx, "No blogs found")
		blankBlogList := schema.BlogList([]schema.Blog{})
		return blankBlogList.ToPaginatedListResp(constants.ApiV1, pageReq.Page, pageReq.PageSize, 0), nil
	}

	// Prepare the paginated response
	blogListResp := make([]resp.BlogPublicResp, len(blogs))
	for i, blog := range blogs {
		blogListResp[i] = blog.ToResponsePublic()
	}

	paginatedResponse := &resp.BlogListPaginatedResp{
		Items:      blogListResp,
		Pagination: resp.NewPaginationResp(constants.ApiV1, pageReq.Page, pageReq.PageSize, totalCount),
	}

	s.log.Infof("Successfully fetched %d blogs", len(blogs))
	return paginatedResponse, nil
}

// CreateBlog adds a new blog to the repository.
func (s *blogService) CreateBlog(ctx context.Context, blog *schema.Blog) error {
	if err := validateBlog(blog); err != nil {
		return err
	}

	s.log.Infof("Creating new blog: %+v", blog)
	if err := s.blogRepo.CreateBlog(context.Background(), blog); err != nil {
		s.log.WithError(err).Error("Failed to create blog")
		return fmt.Errorf("could not create blog: %w", err)
	}
	s.log.Info(ctx, "Successfully created blog")
	return nil
}

// UpdateBlog modifies an existing blog in the repository.
func (s *blogService) UpdateBlog(ctx context.Context, blog *schema.Blog) error {
	if err := validateBlog(blog); err != nil {
		return err
	}

	s.log.Infof("Updating blog: %+v", blog)
	if err := s.blogRepo.UpdateBlog(context.Background(), blog); err != nil {
		s.log.WithError(err).Error("Failed to update blog")
		return fmt.Errorf("could not update blog: %w", err)
	}
	s.log.Info(ctx, "Successfully updated blog")
	return nil
}

// DeleteBlog removes a blog from the repository if the authenticated user is the author.
func (s *blogService) DeleteBlog(ctx context.Context, id int64, authUserID uint) error {
	blog, err := s.blogRepo.GetBlogByID(context.Background(), id)
	if err != nil {
		s.log.WithError(err).Error("Failed to retrieve blog for deletion")
		return fmt.Errorf("could not find blog: %w", err)
	}

	if blog.AuthorID != authUserID {
		s.log.Warn(ctx, "Unauthorized attempt to delete blog")
		return errors.New("unauthorized to delete this blog")
	}

	s.log.Infof("Deleting blog with ID: %d", id)
	if err := s.blogRepo.DeleteBlog(context.Background(), id); err != nil {
		s.log.WithError(err).Error("Failed to delete blog")
		return fmt.Errorf("could not delete blog: %w", err)
	}
	s.log.Info(ctx, "Successfully deleted blog")
	return nil
}

// GetBlogById retrieves a blog by its ID.
func (s *blogService) GetBlogById(ctx context.Context, blogId int64) (*schema.Blog, error) {
	s.log.Infof("Fetching blog with ID: %d", blogId)
	blog, err := s.blogRepo.GetBlogByID(context.Background(), blogId)
	if err != nil {
		s.log.Error(ctx, "Failed to fetch blog by ID")
		return nil, fmt.Errorf("could not retrieve blog: %w", err)
	}
	s.log.Infof("Successfully fetched blog: %+v", blog)
	return blog, nil
}

// GetBlogsByAuthorID retrieves blogs by a specific author with pagination.
func (s *blogService) GetBlogsByAuthorID(ctx context.Context, authorID int64, pageReq request.PaginationRequest) (*resp.BlogListPaginatedResp, error) {
	s.log.Infof("Fetching blogs for author ID: %d with pagination: %+v", authorID, pageReq)

	blogs, totalCount, err := s.blogRepo.GetBlogsByAuthorID(ctx, authorID, pageReq)
	if err != nil {
		s.log.WithError(err).Error("Failed to fetch blogs by author from repository")
		return nil, fmt.Errorf("could not retrieve blogs for author: %w", err)
	}

	if totalCount == 0 {
		s.log.Info(ctx, "No blogs found for this author")
		blankBlogList := schema.BlogList([]schema.Blog{})
		return blankBlogList.ToPaginatedListResp(constants.ApiV1, pageReq.Page, pageReq.PageSize, 0), nil
	}

	blogListResp := make([]resp.BlogPublicResp, len(blogs))
	for i, blog := range blogs {
		blogListResp[i] = blog.ToResponsePublic()
	}

	paginatedResponse := &resp.BlogListPaginatedResp{
		Items:      blogListResp,
		Pagination: resp.NewPaginationResp(constants.ApiV1, pageReq.Page, pageReq.PageSize, totalCount),
	}

	s.log.Infof("Successfully fetched %d blogs for author ID: %d", len(blogs), authorID)
	return paginatedResponse, nil
}

// validateBlog checks if the blog data is valid.
func validateBlog(blog *schema.Blog) error {
	if blog.Title == "" {
		return errors.New("blog title cannot be empty")
	}
	if blog.Content == "" {
		return errors.New("blog content cannot be empty")
	}
	return nil
}
