package controllers

import "auth-service/services"

type Controller interface {
	UserController() UserController
}

type controller struct {
	uCtrl UserController
}

// UserController  is a controller for user
func (c *controller) UserController() UserController {
	return c.uCtrl
}

// NewController  returns a new instance of controller
func NewController(svc services.Services) Controller {
	uSvc := svc.UserService()
	return &controller{
		uCtrl: NewUserController(uSvc),
	}
}
