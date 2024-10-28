package services

import (
	"auth-service/config"
	"auth-service/repositories"
)

// Services is an interface  for services
type Services interface {
	UserService() UserService
}

// svc is the concrete  implementation of the Services interface
type svc struct {
	uSvc UserService
}

// UserService  is the method  to get user service
func (s *svc) UserService() UserService {
	return s.uSvc
}

// NewServices function   to create a new instance of Services
func NewServices(repo repositories.Repository, conf config.Configuration) Services {
	// repo and service init
	userRepo := repo.UserRepository()
	uSvc := NewUserService(userRepo, conf)
	return &svc{
		uSvc: uSvc,
	}
}
