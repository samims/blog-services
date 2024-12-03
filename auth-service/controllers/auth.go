package controllers

import (
	"blog-service/models/resp"
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
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	statusCode, err := c.service.Register(ctx, &req)
	if err != nil {
		c.log.Errorf("Error registering user: %v", err)
		http.Error(w, "User  registration failed", statusCode)

		return
	}
	c.log.Info("User registered successfully")

	w.WriteHeader(http.StatusCreated)
}

func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	status, data, err := c.service.Login(ctx, req)
	if err != nil {
		c.log.WithFields(logrus.Fields{
			"email": req.Email,
			"error": err,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Verify handles  JWT token verification
func (c *authController) Verify(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}

// Logout  handles user logout
func (c *authController) Logout(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}

// RefreshToken  refreshes  JWT token refresh
func (c *authController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	panic("implement")
}
