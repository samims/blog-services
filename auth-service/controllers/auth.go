package controllers

import (
	"encoding/json"
	"net/http"

	"auth-service/models"
	"auth-service/services"

	"github.com/sirupsen/logrus"
)

// AuthController  handles user related operations
type AuthController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Verify(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
}

// implement UserController interface
type authController struct {
	service services.UserService
	log     *logrus.Logger
}

// NewAuthController returns a new instance of authController
func NewAuthController(svc services.UserService, l *logrus.Logger) AuthController {
	return &authController{
		service: svc,
		log:     l,
	}
}

// Register  handles user registration
func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	var req models.User
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		c.log.Warn("Failed to decode request body: ", err)
		RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	statusCode, err := c.service.Register(ctx, &req)
	if err != nil {
		c.log.Errorf("Error registering user: %v", err)
		RespondWithJSON(w, statusCode, nil, err.Error())
		return
	}
	c.log.Info("User registered successfully")
	RespondWithJSON(w, http.StatusCreated, nil, "")
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	status, data, err := c.service.Login(ctx, req)
	if err != nil {
		logrus.Warnf("Error logging in user: %v", err)
		RespondWithError(w, status, err.Error())
		return
	}

	RespondWithJSON(w, status, data, "")
}

// Verify handles  JWT token verification
func (c *authController) Verify(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		RespondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
		return
	}

	// call the service to verify the token
	verified, err := c.service.VerifyToken(r.Context(), tokenString)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, http.StatusOK, verified, "")

}

// Logout  handles user logout
func (c *authController) Logout(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}

// RefreshToken  refreshes  JWT token refresh
func (c *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}
