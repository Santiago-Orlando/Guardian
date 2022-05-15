package filestorage

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var URL string = "FILE_STORAGE_URL"
var PORT string = ":" + os.Getenv("PORT_FILESTORAGE_SERVICE")

func File(w http.ResponseWriter, r *http.Request) {

	endpoint := "/file"
	route := "http://" + URL + PORT + endpoint

	cookie := r.Header.Get("Cookie")

	filename := r.URL.Query().Get("filename")
	if filename != "" {
		route += "?filename=" + filename
	}

	req, err := http.NewRequest(r.Method, route, r.Body)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}
	
	req.Header.Add("Cookie", cookie)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(w, "Error, %s", err)
	}

	w.WriteHeader(response.StatusCode)
	w.Write(data)
}
