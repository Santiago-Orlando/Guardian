package lib

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	m "Guardian/authentication/api/models"
)

func JWTValidator(token string) error {

	jwtStructure := &m.JWTStructure{}

	parsedToken, err := jwt.ParseWithClaims(token, jwtStructure, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWTSecret")), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("Token is not valid!")
	}

	return nil
}
