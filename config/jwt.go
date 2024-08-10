package config

import "github.com/golang-jwt/jwt/v5"

var JWT_KEY = []byte("my_secret_key")

type JWTClaim struct {
	Email string
	jwt.RegisteredClaims
}
