package routes

import (
	"net/http"

	m "Guardian/fileStorage/api/middlewares"
	h "Guardian/fileStorage/api/handlers"
)


func Routes() {

	go http.HandleFunc("/uploadMultipart", m.JWTValidator(h.MultipartFileStorage))
	go http.HandleFunc("/uploadSinglepart", m.JWTValidator(h.SinglepartFileStorage))
	go http.HandleFunc("/getFile", m.JWTValidator(h.FileSender))
	
}