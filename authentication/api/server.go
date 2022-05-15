package api

import (
	"net/http"

	"Guardian/authentication/api/routes"
)


func NewServer(port string) *http.Server{
	
	routes.Routes()

	return &http.Server{
		Addr: port,

	}
}