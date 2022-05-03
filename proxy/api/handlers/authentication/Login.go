package handlers

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

	endpoint := "/login"

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

	cookie := response.Cookies()[0]

	token := cookie.Value

	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Value: token,
	})

	w.WriteHeader(response.StatusCode)
}
