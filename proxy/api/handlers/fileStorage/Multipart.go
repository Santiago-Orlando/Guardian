package filestorage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"Guardian/proxy/api/lib"
)

const MAX_UPLOAD_SIZE int64 = 1024 * 1024 * 100 // 100MB

func Multipart(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	multiFile, fileHeader, err := r.FormFile("filename")
	if err != nil {
		lib.ErrorHandler(err, "web")
		http.Error(w, err.Error(), 400)
		return
	}
	defer multiFile.Close()

	file := lib.MultipartToSinglepart(multiFile)

	endpoint := "/file"

	route := "http://" + URL + PORT + endpoint + "?filename=" + fileHeader.Filename

	cookie := r.Header.Get("Cookie")
	

	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(file)
	

	req, err := http.NewRequest(r.Method, route, &buf)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	w.WriteHeader(response.StatusCode)
}
