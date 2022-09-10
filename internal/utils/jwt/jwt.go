package jwt

import (
	jwtConfig "github.com/Gym-Apps/gym-backend/internal/config/jwt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
)

var JWTConfig = middleware.JWTConfig{
	Claims:     &JwtCustomClaims{},
	SigningKey: []byte(jwtConfig.TokenSecret),
}

type JwtCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
}
