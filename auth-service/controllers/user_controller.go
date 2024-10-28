package controllers

import (
	"auth-service/models"
	"auth-service/services"

	"encoding/json"
	"net/http"

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

// NewUserController returns a new instance of UserController
func NewUserController(svc services.UserService) AuthController {
	return &authController{service: svc}
}

// Register  handles user registration
func (c *authController) Register(w http.ResponseWriter, r *http.Request) {
	var req models.User

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.Register(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

// Login controller handles  user login and  returns a JWT token
func (c *authController) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := c.service.Login(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = json.NewEncoder(w).Encode(resp)

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
