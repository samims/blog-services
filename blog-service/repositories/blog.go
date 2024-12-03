package repositories

import (
	"blog-service/logger"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"blog-service/models"
	"blog-service/models/schema"

	"blog-service/models/request"
)

// BlogRepository defines the methods for interacting with the blog data.
type BlogRepository interface {
	CreateBlog(ctx context.Context, blog *schema.Blog) error
	GetAllBlogs(ctx context.Context, pageReq request.PaginationRequest) ([]schema.Blog, int64, error)
	GetBlogsByAuthorID(ctx context.Context, authorId int64, pageReq request.PaginationRequest) ([]schema.Blog, int64, error)

	GetBlogCount(ctx context.Context) (int64, error)

	GetBlogByID(ctx context.Context, blogId int64) (*schema.Blog, error)
	UpdateBlog(ctx context.Context, blog *schema.Blog) error
	DeleteBlog(ctx context.Context, blogId int64) error
}

// blogRepository is a concrete implementation of BlogRepository.
type blogRepository struct {
	db  *sql.DB
	log *logger.AppLogger
}

// NewBlogRepository creates a new instance of BlogRepository.
func NewBlogRepository(db *sql.DB, log *logger.AppLogger) BlogRepository {
	return &blogRepository{
		db:  db,
		log: log,
	}
}

// CreateBlog adds a new blog to the repository.
func (repo *blogRepository) CreateBlog(ctx context.Context, blog *schema.Blog) error {
	repo.log.Infof("Creating new blog: %+v", blog)
	query := `INSERT INTO blogs (title, content, author_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`

	_, err := repo.db.ExecContext(ctx, query, blog.Title, blog.Content, blog.AuthorID, time.Now(), time.Now())
	if err != nil {
		repo.log.Errorf("Failed to create blog: %v", err)
		return fmt.Errorf("creating blog: %w", err)
	}
	return nil
}

// GetAllBlogs retrieves all blogs with pagination.
func (repo *blogRepository) GetAllBlogs(ctx context.Context, pageReq request.PaginationRequest) ([]schema.Blog, int64, error) {
	var blogs = make([]schema.Blog, 0)
	var totalRecords int64 = 0
	var err error
	repo.log.Info(ctx, "Getting all blogs")

	// Counting total records
	countQuery := `SELECT COUNT(*) FROM blogs`
	if err := repo.db.QueryRowContext(ctx, countQuery).Scan(&totalRecords); err != nil {
		repo.log.Errorf("Failed to count total blogs: %v", err)
		return blogs, totalRecords, fmt.Errorf("counting blogs: %w", err)
	}
	repo.log.Debugf("Total blogs count: %d", totalRecords)

	// Fetching paginated blogs
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM blogs LIMIT $1 OFFSET $2`
	rows, err := repo.db.QueryContext(ctx, query, pageReq.PageSize, pageReq.GetOffset())
	if err != nil {
		repo.log.Errorf("Failed to fetch blogs: %v", err)
		return blogs, totalRecords, fmt.Errorf("fetching blogs: %w", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			repo.log.Errorf("Failed to close rows: %v", err)

		}
	}(rows)

	for rows.Next() {
		var blog schema.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			repo.log.Errorf("Failed to scan blog: %v", err)
			return []schema.Blog{}, 0, fmt.Errorf("scanning blog: %w", err)
		}
		blogs = append(blogs, blog)
	}

	if err := rows.Err(); err != nil {
		repo.log.Errorf("Row iteration error: %v", err)
		return []schema.Blog{}, 0, fmt.Errorf("iterating rows: %w", err)
	}

	return blogs, totalRecords, nil
}

// GetBlogsByAuthorID retrieves blogs by a specific author.
func (repo *blogRepository) GetBlogsByAuthorID(ctx context.Context, authorId int64, pageReq request.PaginationRequest) ([]schema.Blog, int64, error) {
	repo.log.Infof("Fetching blogs by author ID: %d", authorId)

	var blogs = make([]schema.Blog, 0)
	var totalRecords int64 = 0
	var err error

	authorBlogCountQuery := `SELECT COUNT(*) FROM blogs WHERE author_id = ?`

	if err := repo.db.QueryRowContext(ctx, authorBlogCountQuery, authorId).Scan(&totalRecords); err != nil {
		repo.log.Errorf("Failed to count total blogs: %v", err)
		return blogs, totalRecords, fmt.Errorf("counting blogs: %w", err)
	}

	repo.log.Debugf("Total blogs count: %d", totalRecords)

	queryStr := `SELECT id, title, content, author_id, created_at, updated_at FROM blogs WHERE author_id = ? LIMIT ? OFFSET ?`
	rows, err := repo.db.QueryContext(ctx, queryStr, authorId, pageReq.PageSize, pageReq.GetOffset())
	if err != nil {
		repo.log.Errorf("Failed to fetch blogs: %v", err)
		return blogs, 0, fmt.Errorf("fetching blogs: %w", err)
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			repo.log.Errorf("Failed to close rows: %v", err)
			return
		}
	}(rows)

	for rows.Next() {
		var blog schema.Blog
		if err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
			repo.log.Errorf("Failed to scan blog: %v", err)
			return blogs, 0, fmt.Errorf("scanning blog: %w", err)
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		repo.log.Errorf("Row iteration error: %v", err)
		return blogs, 0, fmt.Errorf("iterating rows: %w", err)
	}

	return blogs, totalRecords, nil

}

// GetBlogCount retrieves the total number of blogs
func (repo *blogRepository) GetBlogCount(ctx context.Context) (int64, error) {
	var count int64
	query := `SELECT COUNT(*) FROM blogs`
	err := repo.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		repo.log.Errorf("Failed to fetch total blog count: %v", err)
		return 0, fmt.Errorf("fetching total blog count: %w", err)
	}
	return count, nil
}

// GetBlogByID retrieves a blog by its ID.
func (repo *blogRepository) GetBlogByID(ctx context.Context, blogId int64) (*schema.Blog, error) {
	repo.log.Infof("Fetching blog by ID: %d", blogId)
	query := `SELECT id, title, content, author_id, created_at, updated_at FROM blogs WHERE id = ?`
	var blog schema.Blog

	if err := repo.db.QueryRowContext(ctx, query, blogId).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.AuthorID, &blog.CreatedAt, &blog.UpdatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			repo.log.Warnf("Blog not found with ID: %d", blogId)
			return nil, models.ErrBlogNotFound
		}
		repo.log.Errorf("Failed to scan blog: %v", err)
		return nil, fmt.Errorf("fetching blog by ID: %w", err)
	}
	return &blog, nil
}

// UpdateBlog modifies an existing blog in the repository.
func (repo *blogRepository) UpdateBlog(ctx context.Context, blog *schema.Blog) error {
	repo.log.Infof("Updating blog: %+v", blog)
	query := `UPDATE blogs SET title = ?, content = ?, author_id = ?, updated_at = ? WHERE id = ?`

	_, err := repo.db.ExecContext(ctx, query, blog.Title, blog.Content, blog.AuthorID, time.Now(), blog.ID)
	if err != nil {
		repo.log.Errorf("Failed to update blog: %v", err)
		return fmt.Errorf("updating blog: %w", err)
	}
	return nil
}

// DeleteBlog removes a blog from the repository.
func (repo *blogRepository) DeleteBlog(ctx context.Context, blogId int64) error {
	repo.log.Infof("Deleting blog with ID: %d", blogId)
	query := `DELETE FROM blogs WHERE id = ?`

	_, err := repo.db.ExecContext(ctx, query, blogId)
	if err != nil {
		repo.log.Errorf("Failed to delete blog: %v", err)
		return fmt.Errorf("deleting blog: %w", err)
	}
	return nil
}
