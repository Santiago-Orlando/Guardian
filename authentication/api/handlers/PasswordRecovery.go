package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"

	"golang.org/x/crypto/bcrypt"

	"Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
)

func ForgotPassword(w http.ResponseWriter, r *http.Request) {

	email := &m.Email{}

	err := json.NewDecoder(r.Body).Decode(email)
	if err != nil {
		fmt.Fprintf(w, "error -> ", err)
		return
	}

	collections, err := database.GetConnection("users")
	if err != nil {
		fmt.Fprintf(w, "error db -> ", err)
		return
	}

	token, err := lib.JWTGenerator(email.Email)


	filter := bson.D{{"email", email.Email}}
	update := bson.D{{"$set", bson.D{{"recoveryToken", token}} }}

	user := &m.User{}

	result := collections.FindOneAndUpdate(context.TODO(), filter, update).Decode(user)
	if result != nil {
		fmt.Fprintf(w, "Email not found!")
		return
	}

	lib.SendEmail(user.Email, token)

	w.WriteHeader(200)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	data := &m.PasswordRecovery{}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		fmt.Fprintf(w, "error -> ", err)
	}

 	err = lib.JWTValidator(data.Token)
	if err != nil {
		w.WriteHeader(401)
		return
	}

	collections, err := database.GetConnection("users")
	if err != nil {
		fmt.Fprintf(w, "error db -> ", err)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.NewPassword), 15)
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.D{{"recoveryToken", data.Token}}
	update := bson.D{{"$set", bson.D{{"password", string(password)}, {"recoveryToken", ""} }}}

	var user bson.M

	err = collections.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error -> %s", err)
		return
	}

	w.WriteHeader(202)
}
