package routes

import (
	"net/http"

	"Guardian/fileStorage/api/handlers"
	h "Guardian/fileStorage/api/handlers"
	m "Guardian/fileStorage/api/middlewares"

	"github.com/go-openapi/runtime/middleware"
)

func Routes() {

	go http.HandleFunc("/uploadMultipart", m.JWTValidator(h.MultipartFileStorage))
	go http.HandleFunc("/uploadSinglepart", m.JWTValidator(h.SinglepartFileStorage))
	go http.HandleFunc("/getFile", m.JWTValidator(h.FileSender))

	go http.Handle("/docs", middleware.Redoc(middleware.RedocOpts{SpecURL: "/docs/swagger.yaml"}, nil))
	go http.HandleFunc("/docs/swagger.yaml", handlers.YamlServer)

}
