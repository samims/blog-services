package router

import (
	"auth-service/controllers"
	"fmt"
	"net/http"
)

// InitUserRouter  initializes the user router
func InitUserRouter(ctrl controllers.Controller) *http.ServeMux {
	userCtrl := ctrl.UserController()

	loginPath := fmt.Sprintf("%s /login", http.MethodPost)
	registerPath := fmt.Sprintf("%s /register", http.MethodPost)

	router := http.ServeMux{}

	router.HandleFunc(loginPath, userCtrl.Login)
	router.HandleFunc(registerPath, userCtrl.Register)

	return &router

}
