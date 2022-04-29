package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"

	"Guardian/fileStorage/api/lib"
)

const MAX_UPLOAD_SIZE int64 = 1024 * 1024 //* 1024 * 10 // 10GB

func MultipartFileStorage(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}


	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		lib.ErrorHandler(err, "web")
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 10GB in size", 400)
		return
	}


	multiFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		lib.ErrorHandler(err, "web")
		http.Error(w, err.Error(), 400)
		return
	}
	defer multiFile.Close()


	validate := validator.New()
	err = validate.Var(fileHeader.Filename, "required,gte=3,lte=30")
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}


	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)


	unixTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	newFilename := unixTime + filepath.Ext(fileHeader.Filename)
	path := "./fileStorage/uploads/" + newFilename

	
	file := lib.MultipartToSinglepart(multiFile)


	sha256 := lib.HashFile(file)


	data, err := lib.DuplicatePreventor(sha256, userID)
	if err == nil {

		err = lib.FileSaveDB(fileHeader.Filename, data.Sha, data.NewFileName, userID)
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


	err = lib.FileSaveDB(fileHeader.Filename, sha256, newFilename, userID)
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
