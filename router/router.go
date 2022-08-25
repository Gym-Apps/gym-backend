package router

import (
	"github.com/Gym-Apps/gym-backend/internal/app/user"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func setup(tx *gorm.DB) {
	user.UserInit(tx)
}

func Init(router *echo.Echo, tx *gorm.DB) {
	setup(tx)
	userHandler := user.Handler
	router.POST("/login", userHandler.Login)
}
