package repositories

import (
	"auth-service/models"
	"database/sql"
)

// UserRepository is a repository for user data
type UserRepository interface {
	Create(user *models.User) error
	GetByUserEmail(email string) (models.User, error)
}

// userRepository is a concrete implementation of UserRepository
type userRepository struct {
	DB *sql.DB
}

// NewUserRepository  returns a new instance of userRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{DB: db}
}

// Create creates a new user in the database
func (r userRepository) Create(user models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, username,  password) VALUES ($1, $2, $3, $4)", user.ID, user.Username, user.Password)
	return err
}

// GetByUserName returns a user by their username
func (r userRepository) GetByUserName(username string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT id, username FROM users WHERE username = $1 ",
		username,
	).Scan(
		&user.ID,
		&user.Username,
	)
	return user, err
}
