package routes

import (
	"net/http"

	h "Guardian/fileStorage/api/handlers"
	m "Guardian/fileStorage/api/middlewares"
)

func Routes() {

	go http.HandleFunc("/multipart", m.JWTValidator(h.Multipart))
	go http.HandleFunc("/file", m.JWTValidator(h.File))

}
