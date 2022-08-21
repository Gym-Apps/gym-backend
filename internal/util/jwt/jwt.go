package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
