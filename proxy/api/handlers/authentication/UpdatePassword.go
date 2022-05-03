package handlers

import (
	"fmt"
	"net/http"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	endpoint := "/update-password"

	route := URL + PORT + endpoint

	req, err := http.NewRequest(r.Method, route, r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	w.WriteHeader(response.StatusCode)
}
