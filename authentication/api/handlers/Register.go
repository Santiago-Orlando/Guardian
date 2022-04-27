package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "Guardian/authentication/api/database"
	m "Guardian/authentication/api/models"
)

func Register(w http.ResponseWriter, r *http.Request) {

	connection := db.GetConnection()
	

	user := &m.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		fmt.Println(err)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		fmt.Println(err)
	}
	user.Password = string(password)

	_, err = connection.InsertOne(context.Background(), user)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(201)
}
