package lib

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"

	m "Guardian/authentication/api/models"
)


func JWTGenerator(ID string) (string, error) {

	expirationTime := time.Now().Add(336 * time.Hour) // 14 days

	JWT := &m.JWTStructure{
		ID: ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWT)

	JWTKey := []byte(os.Getenv("JWTSecret"))

	signedToken, err := token.SignedString(JWTKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
