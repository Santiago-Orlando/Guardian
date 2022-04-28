package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"

	db "Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
)


func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}


	connection := db.GetConnection()

	user := &m.User{}

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	validate := validator.New()

	err = validate.Struct(user)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(401)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}
	user.Password = string(password)

	_, err = connection.InsertOne(context.Background(), user)
	if err != nil {
		lib.ErrorHandler(err, "database")
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}
