package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTStructure struct {
	ID				string
	jwt.StandardClaims
}