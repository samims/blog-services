package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"auth-service/config"
	"auth-service/constants"
	"auth-service/models"
	"auth-service/repositories"
	"auth-service/utils"

	"golang.org/x/crypto/bcrypt"
)

// UserService is an  interface for user service
type UserService interface {
	Register(ctx context.Context, user *models.User) (int, error)
	Login(ctx context.Context, loginReq models.LoginRequest) (int, models.LoginResponse, error)
	VerifyToken(ctx context.Context, token string) (int, error)
	RefreshToken(ctx context.Context, token string) (int, models.LoginResponse, error)
}

// userService is an implementation of UserService
type userService struct {
	repo repositories.UserRepository
	conf config.Configuration
}

// Register creates a new user after hashing the password
func (u userService) Register(ctx context.Context, user *models.User) (int, error) {

	existingUser, err := u.repo.GetByUserEmail(ctx, user.Email)

	if err != nil {
		log.Printf("error while getting user by email: %v", err)
		return http.StatusBadRequest, err
	}
	log.Printf("existing user %v", existingUser)
	if existingUser.ID != 0 {
		log.Printf("user with email %s already exists", user.Email)
		return http.StatusBadRequest, fmt.Errorf("user with email %s already exists", user.Email)
	}

	// hash the password, because we don't want to store plain text password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.Password = string(hashedPassword)

	// save the user
	err = u.repo.Create(ctx, user)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

// Login service handles business logic  for login
func (u userService) Login(ctx context.Context, loginReq models.LoginRequest) (int, models.LoginResponse, error) {
	user, err := u.repo.GetByUserEmail(ctx, loginReq.Email)
	if err != nil {
		log.Println("user fetch error")
		return http.StatusUnauthorized, models.LoginResponse{}, errors.New(constants.ErrInvalidEmailOrPass)
	}

	log.Printf("db password %s form password %s", user.Password, loginReq.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password)); err != nil {
		log.Println("user bcrypt error ", err.Error())
		return http.StatusInternalServerError, models.LoginResponse{}, errors.New(constants.ErrInvalidEmailOrPass)
	}

	// generate token
	claims := utils.NewTokenClaims(user.Email, time.Now().UTC().Unix())

	secretKey := u.conf.AppConfig().SecretKey()
	expiresAt := time.Now().UTC().Add(time.Hour * 24 * 7).Unix()

	tokenStr, err := utils.GenerateTokenWithCustomClaims(claims, secretKey, expiresAt)
	if err != nil {
		log.Println("error while generating token", err.Error())
		return http.StatusInternalServerError, models.LoginResponse{}, err
	}

	// prepare response
	loginResp := models.LoginResponse{
		AccessToken: tokenStr,
		Email:       user.Email,
		ExpiresAt:   expiresAt,
	}
	log.Printf("user logged in successfully: %s", loginReq.Email)

	return http.StatusOK, loginResp, nil
}

// VerifyToken provides the business logic to verify JWT tokens
func (u userService) VerifyToken(_ context.Context, token string) (int, error) {
	secretKey := u.conf.AppConfig().SecretKey()

	const (
		statusUnauthorized = http.StatusUnauthorized
		errInvalidToken    = "invalid token"
	)

	if err := u.validateToken(token, secretKey); err != nil {
		log.Println("error validating token:", err)
		return statusUnauthorized, err
	}

	return http.StatusOK, nil
}

func (u userService) RefreshToken(_ context.Context, token string) (int, models.LoginResponse, error) {
	// TODO: implement refresh token logic
	panic("")
}

// validateToken checks if the provided token is valid.
func (u userService) validateToken(token, secretKey string) error {
	if _, err := utils.ValidateToken(token, secretKey); err != nil {
		return errors.New("invalid token")
	}
	return nil
}

// NewUserService returns a new instance of the service
func NewUserService(repo repositories.UserRepository, conf config.Configuration) UserService {
	return &userService{
		repo: repo,
		conf: conf,
	}
}
