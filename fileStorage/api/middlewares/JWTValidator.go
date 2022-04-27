package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	m "Guardian/fileStorage/api/models"
)

func JWTValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token := cookie.Value
		jwtStructure := &m.JWTStructure{}

		parsedToken, err := jwt.ParseWithClaims(token, jwtStructure, func(tkn *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSecret")), nil // os
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "%v", err)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("entre acá", err)
			return
		}

		if !parsedToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "%v", err)
			return
		}

		ctxValues := context.WithValue(r.Context(), "id", jwtStructure.ID)
		r = r.WithContext(ctxValues)


		next(w, r)
	})
}
