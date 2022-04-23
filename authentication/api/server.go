package api

import (
	"net/http"

	"Guardian/authentication/api/routes"
)


func NewServer(addr string) *http.Server{

	routes.Routes()

	return &http.Server{
		Addr: addr,

	}
}