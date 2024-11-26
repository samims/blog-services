package repositories

import (
	"database/sql"
	"fmt"
)

// Repository is an interface for all repositories
type Repository interface {
	UserRepository() UserRepository
}

// repo  is a concrete  implementation of Repository
type repo struct {
	userRepository UserRepository
}

// UserRepository implements Repository.
func (r *repo) UserRepository() UserRepository {
	return r.userRepository
}

// NewRepository returns a new instance of Repository.
func NewRepository(db *sql.DB) (Repository, error) {
	if db == nil {
		return nil, fmt.Errorf("db connection cannot be nil")
	}
	return &repo{
		userRepository: NewUserRepository(db),
	}, nil
}
