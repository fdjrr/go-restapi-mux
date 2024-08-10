package middlewares

import (
	"github/fdjrr/go-restapi-mux/config"
	"github/fdjrr/go-restapi-mux/helpers"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

var ResponseError = helpers.ResponseError

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")

		if err != nil {
			switch err {
			case http.ErrNoCookie:
				ResponseError(w, http.StatusUnauthorized, "Unauthorized")
				return

			default:
				ResponseError(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		tokenString := c.Value

		claims := &config.JWTClaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			switch err {
			case jwt.ErrSignatureInvalid:
				ResponseError(w, http.StatusUnauthorized, "Unauthorized")
				return

			default:
				ResponseError(w, http.StatusBadRequest, err.Error())
				return
			}
		}

		if !token.Valid {
			ResponseError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}
