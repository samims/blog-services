package repositories

import (
	"blog-service/models/schema"
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
	Create(ctx context.Context, blog *schema.Blog) error
	Update(ctx context.Context, blog *schema.Blog) error
	Delete(ctx context.Context, blog *schema.Blog) error
	GetByID(ctx context.Context, blogId int64) (*schema.Blog, error)
	List(ctx context.Context, blog *[]schema.Blog, limit int, offset int) (*[]schema.Blog, error)
	GetByAuthor(ctx context.Context, authorID uint, limit, offset int) (*[]schema.Blog, error)
}

// blogRepository is a concrete implementation of BlogRepository
type blogRepository struct {
	db  *sql.DB
	log *logrus.Logger
}

func (r *blogRepository) GetByAuthor(ctx context.Context, authorID uint, limit, offset int) (*[]schema.Blog, error) {
	//TODO implement me
	panic("implement me")
}

// NewBlogRepository returns a new instance of BlogRepository
func NewBlogRepository(db *sql.DB, log *logrus.Logger) BlogRepository {
	return &blogRepository{
		db:  db,
		log: log,
	}
}

// Create inserts  a new blog into the database
func (r *blogRepository) Create(ctx context.Context, blog *schema.Blog) error {
	// Create a new transaction
	if blog == nil {
		return ErrInvalidBlog
	}

	//query := `
	//INSERT INTO blogs (title, content, author_id, created_at, updated_at)
	//VALUES ($1, $2, $3, $4, $5)
	//RETURNING id`

	// set the timestamps
	now := time.Now().UTC()
	blog.CreatedAt = now
	blog.UpdatedAt = now

	return nil

}

// Update ...
func (r *blogRepository) Update(ctx context.Context, blog *schema.Blog) error {
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
func (r *blogRepository) Delete(ctx context.Context, blog *schema.Blog) error {
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
func (r *blogRepository) GetByID(ctx context.Context, blogId int64) (*schema.Blog, error) {
	var blog schema.Blog

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
func (r *blogRepository) List(ctx context.Context, blog *[]schema.Blog, limit int, offset int) (*[]schema.Blog, error) {
	var blogs []schema.Blog

	queryStr := `SELECT id, title, content, author, created_at from blogs ORDER BY created_at  DESC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(queryStr, limit, offset)

	if err != nil {
		r.log.Errorf("Error listing blogs: %v", err)
		return nil, err
	}

	defer rows.Close()
	// scan  rows and  append to blog slice
	for rows.Next() {
		var blog schema.Blog
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
