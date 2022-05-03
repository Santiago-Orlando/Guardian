package docs

import (
	"io/ioutil"
	"net/http"
	"os"
)

func SwaggerServer(w http.ResponseWriter, r *http.Request) {

	file, err := os.Open("./proxy/swagger.yaml")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(data)
}
