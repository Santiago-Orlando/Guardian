package lib

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	m "Guardian/authentication/api/models"
)

func ErrorHandler(err error, category string) {

	data := m.Err{
		Err:      err.Error(),
		Category: category,
	}
	buf := bytes.NewBuffer(nil)

	_ = json.NewEncoder(buf).Encode(data)

	url := "http://ERROR_LOGGER_URL:" + os.Getenv("PORT_ERROR_SERVICE") + "/errors"
	contentType := "application/json"

	http.Post(url, contentType, buf)
}
