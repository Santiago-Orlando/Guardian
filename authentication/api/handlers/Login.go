package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

	db "Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		w.WriteHeader(405)
	}

	// Create database connection
	connection := db.GetConnection()

	userData := &m.UserLogin{}
	user := &m.User{}

	err := json.NewDecoder(r.Body).Decode(userData)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	validate := validator.New()
	
	err = validate.Struct(userData)
	if err != nil {
		lib.ErrorHandler(err, "authentication")
		w.WriteHeader(406)
		return
	}
	
	// Find User in database
	err = connection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: userData.Email}}).Decode(user)
	if err != nil {
		lib.ErrorHandler(err, "database")
		w.WriteHeader(500)
		return
	}

	// Compare the password that was send with the one in the DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	//Generate Token
	token, err := lib.JWTGenerator(user.ID)
	if err != nil {
		lib.ErrorHandler(err, "system")
		w.WriteHeader(500)
		return
	}

	// Send token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "jwt",
		Value: token,
	})

}
