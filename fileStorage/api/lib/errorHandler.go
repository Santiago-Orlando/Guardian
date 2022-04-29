package lib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	m "Guardian/fileStorage/api/models"
)

func ErrorHandler(err error, category string) {

	fmt.Println(err)

	data := m.Err{
		Err:      err.Error(),
		Category: category,
	}
	buf := bytes.NewBuffer(nil)

	_ = json.NewEncoder(buf).Encode(data)

	url := "http://localhost:" + os.Getenv("PORT_ERROR_SERVICE") + "/errors"
	contentType := "application/json"

	http.Post(url, contentType, buf)
}
