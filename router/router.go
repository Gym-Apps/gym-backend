package router

import (
	"github.com/Gym-Apps/gym-backend/internal/app/user"
	"github.com/Gym-Apps/gym-backend/internal/middlewares"
	"github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func setup(tx *gorm.DB) {
	user.UserInit(tx)
}

func Init(router *echo.Echo, tx *gorm.DB) {
	setup(tx)
	userHandler := user.Handler
	router.POST("/login", userHandler.Login)

	auth := router.Group("")
	auth.Use(middleware.JWTWithConfig(jwt.JWTConfig))
	auth.Use(middlewares.VerifyToken)

	auth.POST("/reset/password", userHandler.ResetPassword)
}
