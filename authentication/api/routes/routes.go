package routes

import (
	"Guardian/authentication/api/handlers"
	"net/http"
)

func Routes() {

	go http.HandleFunc("/register", handlers.Register)
	go http.HandleFunc("/login", handlers.Login)
	go http.HandleFunc("/forgotPassword", handlers.ForgotPassword)
	go http.HandleFunc("/updatePassword", handlers.UpdatePassword)

}
