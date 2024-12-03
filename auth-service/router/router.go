package router

import (
	"fmt"
	"net/http"

	"auth-service/controllers"
)

// InitUserRouter  initializes the user router
func InitUserRouter(ctrl controllers.Controller) *http.ServeMux {
	userCtrl := ctrl.AuthController()

	loginPath := fmt.Sprintf("%s /login", http.MethodPost)
	registerPath := fmt.Sprintf("%s /register", http.MethodPost)

	router := http.ServeMux{}

	router.HandleFunc(registerPath, userCtrl.Register)
	router.HandleFunc(loginPath, userCtrl.Login)

	return &router

}
