package router

import (
	"github.com/Gym-Apps/gym-backend/internal/app/user"
	"github.com/labstack/echo/v4"
)

func Init(router *echo.Echo) {
	userHandler := user.Handler
	router.GET("/login", userHandler.Login)
}
