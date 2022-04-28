package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"Guardian/fileStorage/api/lib"
	m "Guardian/fileStorage/api/models"

	"github.com/go-playground/validator/v10"
)

func FileSender(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.WriteHeader(405)
	}


	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)

	filename := m.RequestFile{}

	err := json.NewDecoder(r.Body).Decode(&filename)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}


	validate := validator.New()
	err = validate.Struct(filename)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}

	serverName := lib.GetFileFromDB(userID, filename.Filename)

	file, err := os.Open("./uploads/" + serverName)
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
