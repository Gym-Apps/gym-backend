package router

import (
	"github.com/Gym-Apps/gym-backend/internal/app/user"
	"github.com/Gym-Apps/gym-backend/internal/middlewares"
	"github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Init(router *echo.Echo, tx *gorm.DB) {
	userHandler := user.UserInit(tx)
	router.POST("/login", userHandler.Login)

	auth := router.Group("")
	auth.Use(middleware.JWTWithConfig(jwt.JWTConfig))
	auth.Use(middlewares.VerifyToken)

	auth.POST("/reset/password", userHandler.ResetPassword)
}
