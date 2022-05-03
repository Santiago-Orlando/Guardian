package routes

import (
	"net/http"

	ha "Guardian/proxy/api/handlers/authentication"
	hd "Guardian/proxy/api/handlers/docs"
	hf "Guardian/proxy/api/handlers/fileStorage"

	"github.com/go-openapi/runtime/middleware"
)


func Routes() {

	go http.HandleFunc("/register", ha.Register)
	go http.HandleFunc("/login", ha.Login)
	go http.HandleFunc("/forgot-password", ha.ForgotPassword)
	go http.HandleFunc("/update-password", ha.UpdatePassword)


	go http.HandleFunc("/file", hf.File)
	go http.HandleFunc("/multipart", hf.Multipart)


	go http.Handle("/docs", middleware.Redoc(middleware.RedocOpts{SpecURL: "/docs/swagger.yaml"}, nil))
	go http.HandleFunc("/docs/swagger.yaml", hd.SwaggerServer)
}