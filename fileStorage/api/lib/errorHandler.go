package lib

import (
	"bytes"
	"encoding/json"
	"net/http"

	m "Guardian/fileStorage/api/models"
)

func ErrorHandler(err error, category string) {

	data := m.Err{
		Err: err.Error(),
		Category: category,
	}
	buf := bytes.NewBuffer(nil)

	_ = json.NewEncoder(buf).Encode(data)
	
	url := "http://localhost:3003/errors"
	contentType := "application/json"

	http.Post(url, contentType, buf)
}
