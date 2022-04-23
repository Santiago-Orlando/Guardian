package lib

import (
	m "Guardian/authentication/api/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

func JWTValidator(token string) error {
	
	jwtStructure := &m.JWTStructure{}

	parsedToken, err := jwt.ParseWithClaims(token, jwtStructure, func(tkn *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return errors.New("Token is not valid!")
	}

	return nil
}