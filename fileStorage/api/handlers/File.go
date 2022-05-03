package handlers

import (
	"net/http"

	s "Guardian/fileStorage/api/services"
)

func File(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		s.SendFile(w, r)
		return
	}

	if r.Method == "POST" {
		s.StorageFile(w, r)
		return
	}

	if r.Method == "DELETE" {
		s.DeleteFile(w, r)
		return
	}

	w.WriteHeader(405)
}
