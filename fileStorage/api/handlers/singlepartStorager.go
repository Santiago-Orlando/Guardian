package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"Guardian/fileStorage/api/lib"

	"github.com/go-playground/validator/v10"
)

func SinglepartFileStorage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	filename := r.URL.Query().Get("filename")
	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)


	validate := validator.New()
	err := validate.Var(filename, "required,gte=3,lte=30")
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}
	

	unixTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	newFilename := unixTime + filepath.Ext(filename)
	path := "./uploads/" + newFilename


	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}


	sha256 := lib.HashFile(file)

	data, err := lib.DuplicatePreventor(sha256, userID)
	if err == nil {

		err = lib.FileSaveDB(filename, data.Sha, data.NewFileName, userID)
		if err != nil {
			lib.ErrorHandler(err, "database")
			w.WriteHeader(500)
			return
		}

		w.WriteHeader(202)
		fmt.Fprintf(w, "Upload successful")
		return
	}

	if err == http.ErrBodyNotAllowed {
		lib.ErrorHandler(err, "web")
		w.WriteHeader(400)
		return
	}

	err = lib.FileSaveDB(filename, sha256, newFilename, userID)
	if err != nil {
		lib.ErrorHandler(err, "database")
		w.WriteHeader(500)
		return
	}

	err = lib.FileSaver(file, path)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}
