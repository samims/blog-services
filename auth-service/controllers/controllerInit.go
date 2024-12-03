package controllers

import (
	"auth-service/services"
	"github.com/sirupsen/logrus"
)

type Controller interface {
	AuthController() AuthController
}

type controller struct {
	authCtrl AuthController
}

// AuthController ...
func (c *controller) AuthController() AuthController {
	return c.authCtrl
}

// NewController  returns a new instance of controller
func NewController(svc services.Services, l *logrus.Logger) Controller {
	uSvc := svc.UserService()
	return &controller{
		authCtrl: NewAuthController(uSvc, l),
	}
}
