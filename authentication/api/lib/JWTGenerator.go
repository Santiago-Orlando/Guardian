package lib

import (
	"time"

	m "Guardian/authentication/api/models"

	"github.com/dgrijalva/jwt-go"
)


func JWTGenerator(email string) (string, error) {

	expirationTime := time.Now()//.Add(336 * time.Hour) // 14 days

	JWT := &m.JWTStructure{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWT)

	JWTKey := []byte("secret") // os 

	signedToken, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
