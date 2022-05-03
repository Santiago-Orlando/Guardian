package api

import (
	"net/http"

	"Guardian/proxy/api/routes"
)

func NewServer(port string) *http.Server {

	routes.Routes()

	return &http.Server{
		Addr: port,
	}
}
