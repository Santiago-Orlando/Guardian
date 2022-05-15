package handlers

import (
	"net/http"

	"Guardian/proxy/api/lib"
)

func Register(w http.ResponseWriter, r *http.Request) {

	endpoint := "/register"

	route := "http://" + URL + PORT + endpoint

	req, err := http.NewRequest(r.Method, route, r.Body)
	if err != nil {
		lib.ErrorHandler(err, "system")
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		lib.ErrorHandler(err, "web")
	}

	w.WriteHeader(response.StatusCode)
}
