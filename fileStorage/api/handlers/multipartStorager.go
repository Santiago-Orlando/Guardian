package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"Guardian/fileStorage/api/lib"
)

const MAX_UPLOAD_SIZE int64 = 1024 * 1024 //* 1024 * 10 // 10GB

func MultipartFileStorage(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		fmt.Println("eer", err)
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 10GB in size", http.StatusBadRequest)
		return
	}

	multiFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer multiFile.Close()

	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)

	unixTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	newFilename := unixTime + filepath.Ext(fileHeader.Filename)
	path := "./uploads/" + newFilename

	
	file := lib.MultipartToSinglepart(multiFile)

	sha256 := lib.HashFile(file)


	data, err := lib.DuplicatePreventor(sha256, userID)
	if err == nil {

		err = lib.FileSaveDB(fileHeader.Filename, data.Sha, data.NewFileName, userID)
		if err != nil {
			fmt.Println(err)
			return
		}

		w.WriteHeader(202)
		fmt.Fprintf(w, "Upload successful")
		return
	}

	if err == http.ErrBodyNotAllowed {
		w.WriteHeader(400)
		return
	}

	err = lib.FileSaveDB(fileHeader.Filename, sha256, newFilename, userID)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = lib.FileSaver(file, path)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.WriteHeader(201)
}
