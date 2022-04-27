package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"Guardian/fileStorage/api/lib"
	m "Guardian/fileStorage/api/models"
)


func FileSender(w http.ResponseWriter, r *http.Request){

	ctx := r.Context()
	userID, _ := ctx.Value("id").(string)

	filename := m.RequestFile{}

	err := json.NewDecoder(r.Body).Decode(&filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	serverName := lib.GetFileFromDB(userID, filename.Filename)

	file, err := os.Open("./uploads/" + serverName)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(data)
}