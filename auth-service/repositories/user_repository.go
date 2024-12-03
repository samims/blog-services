package repositories

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"auth-service/models"
)

// UserRepository is a repository for user data
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByUserEmail(ctx context.Context, email string) (models.User, error)
}

// userRepository is a concrete implementation of UserRepository
type userRepository struct {
	db *sql.DB
}

// NewUserRepository  returns a new instance of userRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

// Create inserts a new user into the database
func (r userRepository) Create(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Password, user.FirstName, user.LastName).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

// GetByUserEmail retrieves a user by email from the database. Returns a models.User and an error if any occurs.
func (r userRepository) GetByUserEmail(ctx context.Context, email string) (models.User, error) {
	log.Println("getting user by email ", email)
	user := models.User{}
	queryStr := `SELECT id, email, password, first_name, last_name FROM users WHERE email = $1`

	err := r.db.QueryRowContext(ctx, queryStr, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.FirstName,
		&user.LastName,
	)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		log.Printf("Error retrieving user by email %s: %v", email, err)
		return models.User{}, err
	}

	return user, nil
}
