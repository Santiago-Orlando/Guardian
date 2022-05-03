package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
)


func ForgotPassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}


	email := &m.Email{}

	err := json.NewDecoder(r.Body).Decode(email)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	validate := validator.New()

	err = validate.Struct(email)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}

	collections := database.GetConnection()
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	token, err := lib.JWTGenerator(email.Email)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	toUpdate := bson.D{primitive.E{Key: "recoveryToken", Value: token}}
	filter := bson.D{primitive.E{Key: "email", Value: email.Email}}
	update := bson.D{primitive.E{Key:"$set", Value: toUpdate}}

	user := &m.User{}

	err = collections.FindOneAndUpdate(context.TODO(), filter, update).Decode(user)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(404)
		return
	}

	lib.SendEmail(user.Email, token)

	w.WriteHeader(200)
}
