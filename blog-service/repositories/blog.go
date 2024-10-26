package repositories

import (
	"blog-service/models"
	"errors"
	"time"

	"context"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	ErrBlogNotFound = errors.New("blog not found")

	ErrInvalidBlog = errors.New("invalid blog data")

	ErrDBOperation = errors.New("database operation failed")
)

// BlogRepository  is a repository for blog
type BlogRepository interface {
	Create(ctx context.Context, blog *models.Blog) error
	Update(ctx context.Context, blog *models.Blog) error
	Delete(ctx context.Context, blog *models.Blog) error
	GetByID(ctx context.Context, blogId int64) (*models.Blog, error)
	List(ctx context.Context, blog *[]models.Blog, limit int, offset int) (*[]models.Blog, error)
	GetByAuthor(ctx context.Context, authorID uint, limit, offset int) (*[]models.Blog, error)
}

// blogRepository is a concrete implementation of BlogRepository
type blogRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

// NewBlogRepository returns a new instance of BlogRepository
func NewBlogRepository(db *sql.DB, log *logrus.Logger) BlogRepository {
	return &blogRepository{
		db:  db,
		log: log,
	}
}

// Create creates  a new blog using raw SQL .Exec()
func (r *blogRepository) Create(ctx context.Context, blog *models.Blog) error {
	res, err := r.db.Exec(
		"INSERT INTO blogs (title, content, created_at, updated_at) VALUES  (?, ?, ?, ?)",
		blog.Title,
		blog.Content,
		blog.CreatedAt,
		blog.UpdatedAt,
	)
	if err != nil {
		r.log.Errorf("Error creating blog: %v", err)
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		r.log.Errorf("Error getting last insert id: %v", err)
		return err
	}
	blog.ID = uint(id)
	return nil
}

// Update ...
func (r *blogRepository) Update(ctx context.Context, blog *models.Blog) error {
	_, err := r.db.Exec(
		"UPDATE blogs SET title = ?, content = ?, updated_at = ? WHERE id = ?",
		blog.Title,
		blog.Content,
		blog.UpdatedAt,
		blog.ID,
	)
	if err != nil {
		r.log.Errorf("Error updating blog: %v", err)
		return err
	}
	return nil

}

// Delete fetches  a blog by ID and deletes it
func (r *blogRepository) Delete(ctx context.Context, blog *models.Blog) error {
	query := "DELETE FROM blogs WHERE id = $1"

	// execute  the query
	result, err := r.db.Exec(query, blog.ID)
	if err != nil {
		return fmt.Errorf("failed to delete blog %w", err)
	}

	// check the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected %w", err)
	}
	//  if no rows were affected meaning  the blog was not found
	if rowsAffected == 0 {
		return fmt.Errorf("blog not found")
	}

	return nil

}

// GetByID fetches  a blog by its ID.
func (r *blogRepository) GetByID(ctx context.Context, blogId int64) (*models.Blog, error) {
	var blog models.Blog

	// Prepare the query
	query := "SELECT id, title, content, created_at, updated_at FROM blogs WHERE id = $1"

	// Execute the query
	err := r.db.QueryRow(query, blogId).Scan(&blog.ID, &blog.Title, &blog.Content, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // Return a specific error if no rows are found
			return nil, fmt.Errorf("no blog found with id %d", blogId)
		}
		// Return the error encountered
		return nil, fmt.Errorf("failed to get blog: %w", err)
	}

	return &blog, nil
}

// List fetches  all blogs
func (r *blogRepository) List(ctx context.Context, blog *[]models.Blog, limit int, offset int) (*[]models.Blog, error) {
	var blogs []models.Blog

	queryStr := `SELECT id, title, content, author, created_at from blogs ORDER BY created_at  DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(queryStr, limit, offset)

	if err != nil {
		r.log.Errorf("Error listing blogs: %v", err)
		return nil, err
	}

	defer rows.Close()
	// scan  rows and  append to blog slice
	for rows.Next() {
		var blog models.Blog
		err = rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Content,
			&blog.Author,
			&blog.CreatedAt,
		)
		if err != nil {
			r.log.Errorf("Error scanning blog: %v", err)
			return nil, err
		}
		blogs = append(blogs, blog)
	}
	return &blogs, nil

}
