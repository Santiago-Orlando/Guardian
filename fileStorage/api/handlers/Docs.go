package handlers

import (
	"io/ioutil"
	"net/http"
	"os"

	"Guardian/fileStorage/api/lib"
)

func YamlServer(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("./fileStorage/swagger.yaml")
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
