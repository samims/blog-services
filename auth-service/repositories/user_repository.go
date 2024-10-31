package repositories

import (
	"database/sql"

	"auth-service/models"
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
func (r userRepository) Create(user *models.User) error {
	row := r.DB.QueryRow(
		`INSERT INTO users (email, password, first_name, last_name) 
				VALUES ($1, $2, $3, $4)
				RETURNING id`,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
	).Scan(&user.ID)
	if row == nil {
		return row
	}
	return nil
}

// GetByUserEmail returns a user by their username
func (r userRepository) GetByUserEmail(email string) (models.User, error) {
	var user models.User
	err := r.DB.QueryRow(
		"SELECT id, first_name username FROM users WHERE email = $1 ",
		email,
	).Scan(
		&user.ID,
	)
	return user, err
}
