package handlers

import (
	"Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		w.WriteHeader(405)
	}

	data := &m.PasswordRecovery{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	validate := validator.New()

	err = validate.Struct(data)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}

	err = lib.JWTValidator(data.Token)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(401)
		return
	}

	collections := database.GetConnection()
	if err != nil {
		lib.ErrorHandler(err, "database")
		w.WriteHeader(500)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 15)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	toUpdate := bson.D{primitive.E{Key: "password", Value: string(password)}, {Key: "recoveryToken", Value: ""}}
	filter := bson.D{primitive.E{Key: "recoveryToken", Value: data.Token}}
	update := bson.D{primitive.E{Key: "$set", Value: toUpdate}}

	var user bson.M

	err = collections.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		lib.ErrorHandler(err, "database")
		w.WriteHeader(500)
		return
	}

}
