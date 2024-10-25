package controllers

import (
	"auth-service/models"
	"auth-service/services"
	"encoding/json"
	"net/http"
)

// UserController  handles user related operations
type UserController interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
}

// implement UserController interface
type userController struct {
	service services.UserService
}

// NewUserController returns a new instance of UserController
func NewUserController(svc services.UserService) UserController {
	return &userController{service: svc}
}

func (c *userController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c *userController) Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err = c.service.Login(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = json.NewEncoder(w).Encode(user)

}
