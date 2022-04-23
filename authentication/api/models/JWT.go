package models

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTStructure struct {
	Email				string
	jwt.StandardClaims
}