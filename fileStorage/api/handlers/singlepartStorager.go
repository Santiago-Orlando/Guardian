package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"Guardian/fileStorage/api/lib"
)

func SinglepartFileStorage(w http.ResponseWriter, r *http.Request) {

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)

	filename := r.URL.Query().Get("filename")
	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)

	unixTime := strconv.FormatInt(time.Now().UnixNano(), 10)
	newFilename := unixTime + filepath.Ext(filename)
	path := "./uploads/" + newFilename

	file, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	sha256 := lib.HashFile(file)

	data, err := lib.DuplicatePreventor(sha256, userID)
	if err == nil {

		err = lib.FileSaveDB(filename, data.Sha, data.NewFileName, userID)
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

	err = lib.FileSaveDB(filename, sha256, newFilename, userID)
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