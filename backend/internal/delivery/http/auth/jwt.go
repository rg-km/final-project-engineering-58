package auth_http_handler

import (
	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

type Claims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	jwt.StandardClaims
}
