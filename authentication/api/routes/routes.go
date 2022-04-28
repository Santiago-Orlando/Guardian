package routes

import (
	"net/http"

	"Guardian/authentication/api/handlers"

	"github.com/go-openapi/runtime/middleware"
)

func Routes() {

	go http.HandleFunc("/register", handlers.Register)
	go http.HandleFunc("/login", handlers.Login)
	go http.HandleFunc("/forgotPassword", handlers.ForgotPassword)
	go http.HandleFunc("/updatePassword", handlers.UpdatePassword)

	go http.Handle("/docs", middleware.Redoc(middleware.RedocOpts{SpecURL: "/docs/swagger.yaml"}, nil))
	go http.HandleFunc("/docs/swagger.yaml", handlers.YamlServer)

}
