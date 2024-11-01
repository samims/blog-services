package services

import (
	"context"
	"errors"
	"time"

	"auth-service/config"
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/utils"

	"golang.org/x/crypto/bcrypt"
)

// UserService is an  interface for user service
type UserService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, loginReq models.LoginRequest) (models.LoginResponse, error)
	Verify(ctx context.Context, verify models.VerifyRequest) error
}

// userService is an implementation of UserService
type userService struct {
	repo repositories.UserRepository
	conf config.Configuration
}

// Register creates a new user after hashing the password
func (u userService) Register(ctx context.Context, user *models.User) error {
	// hash the password, because we don't want to store plain text password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	// save the user
	err = u.repo.Create(ctx, user)
	return err
}

// Login service handles business logic  for login
func (u userService) Login(ctx context.Context, loginReq models.LoginRequest) (models.LoginResponse, error) {

	user, err := u.repo.GetByUserEmail(ctx, loginReq.Email)
	if err != nil {
		return models.LoginResponse{}, errors.New("username or password error")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginReq.Password), []byte(user.Password)); err != nil {
		return models.LoginResponse{}, errors.New("username or password error")
	}

	// generate token
	claims := utils.NewTokenClaims(user.Email, time.Now().UTC().Unix())

	secretKey := u.conf.AppConfig().SecretKey()
	expiresAt := time.Now().UTC().Add(time.Hour * 24 * 7).Unix()

	tokenStr, err := utils.GenerateTokenWithCustomClaims(claims, secretKey, expiresAt)
	if err != nil {
		return models.LoginResponse{}, err
	}

	// prepare response
	loginResp := models.LoginResponse{
		AccessToken: tokenStr,
		Email:       user.Email,
		ExpiresAt:   expiresAt,
	}

	return loginResp, nil
}

// Verify service  handles business logic for verify it can  be used to verify jwt  token
func (u userService) Verify(_ context.Context, req models.VerifyRequest) error {
	// verify the token
	panic("implement me")
}

// NewUserService returns a new instance of the service
func NewUserService(repo repositories.UserRepository, conf config.Configuration) UserService {
	return &userService{
		repo: repo,
		conf: conf,
	}
}
