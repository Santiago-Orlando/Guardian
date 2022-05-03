package filestorage

import (
	"fmt"
	"net/http"
)

func Multipart(w http.ResponseWriter, r *http.Request) {

	endpoint := "/multipart"

	route := URL + PORT + endpoint

	content := r.Header.Get("Content-Type")
	cookie := r.Header.Get("Cookie")

	req, err := http.NewRequest(r.Method, route, r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	req.Header.Add("Content-Type", content)
	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	w.WriteHeader(response.StatusCode)
}
