package routes

import (
	"net/http"

	"Guardian/authentication/api/handlers"

)

func Routes() {

	go http.HandleFunc("/register", handlers.Register)
	go http.HandleFunc("/login", handlers.Login)
	go http.HandleFunc("/forgot-password", handlers.ForgotPassword)
	go http.HandleFunc("/update-password", handlers.UpdatePassword)

}
