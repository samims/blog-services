package services

import (
	"auth-service/models"
	"auth-service/repositories"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user models.User) error
	Login(user models.User) (models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

// Register creates a new user after hashing the password
func (u userService) Register(user models.User) error {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// save the user
	err = u.repo.Create(user)
	return err
}

func (u userService) Login(user models.User) (models.User, error) {
	user, err := u.repo.GetByUserName(user.Username)
	// need to encrypt
	if err != nil {
		return models.User{}, errors.New("username or password error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(user.Password)); err != nil {
		return models.User{}, errors.New("username or password error")
	}
	return user, nil
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}
