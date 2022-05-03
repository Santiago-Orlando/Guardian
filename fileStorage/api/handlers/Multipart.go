package handlers

import (
	"net/http"

	s "Guardian/fileStorage/api/services"
)

func Multipart(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		s.StorageMultipart(w, r)
		return
	}

	w.WriteHeader(405)

}
