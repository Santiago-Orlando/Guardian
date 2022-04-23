package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"

	db "Guardian/authentication/api/database"
	"Guardian/authentication/api/lib"
	m "Guardian/authentication/api/models"
)

func Login(w http.ResponseWriter, r *http.Request) {

	// Create database connection
	connection, err := db.GetConnection("users")
	if err != nil {
		fmt.Println(err)
	}

	userData := &m.UserLogin{}
	user := &m.User{}

	err = json.NewDecoder(r.Body).Decode(userData)
	if err != nil {
		fmt.Println(err)
	}

	// Find User in database
	err = connection.FindOne(context.TODO(), bson.D{{"email", userData.Email}}).Decode(user)
	if err != nil {
		fmt.Println(err)
	}

	// Compare the password that was send with the one in the DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password))
	if err != nil {
		fmt.Println(err)
	}

	//Generate Token
	token, err := lib.JWTGenerator(user.Email)

	// Send token as cookie
	http.SetCookie(w, &http.Cookie{
		Name:		"jwt",
		Value:		token,
	})

}
