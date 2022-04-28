package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"Guardian/fileStorage/api/lib"
	m "Guardian/fileStorage/api/models"
)

func JWTValidator(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("jwt")
		if err != nil {
			if err == http.ErrNoCookie {
				lib.ErrorHandler(err, "authentication")
				w.WriteHeader(401)
				return
			}
			lib.ErrorHandler(err, "web")
			w.WriteHeader(400)
			return
		}
		token := cookie.Value
		jwtStructure := &m.JWTStructure{}

		parsedToken, err := jwt.ParseWithClaims(token, jwtStructure, func(tkn *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWTSecret")), nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(401)
				fmt.Fprintf(w, "%v", err)
				return
			}
			w.WriteHeader(400)
			return
		}

		if !parsedToken.Valid {
			w.WriteHeader(401)
			fmt.Fprintf(w, "%v", err)
			return
		}

		ctxValues := context.WithValue(r.Context(), "id", jwtStructure.ID)
		r = r.WithContext(ctxValues)

		next(w, r)
	})
}
