package routes

import (
	"net/http"
	
	"Guardian/authentication/api/handlers"
)

func Routes() {

	go http.HandleFunc("/register", handlers.Register)
	go http.HandleFunc("/login", handlers.Login)
	go http.HandleFunc("/forgotPassword", handlers.ForgotPassword)
	go http.HandleFunc("/updatePassword", handlers.UpdatePassword)

}
