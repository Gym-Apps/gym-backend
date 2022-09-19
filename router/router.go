package router

import (
	"github.com/Gym-Apps/gym-backend/internal/app/user"
	"github.com/Gym-Apps/gym-backend/internal/middlewares"
	"github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"gorm.io/gorm"
)

func Init(router *echo.Echo, tx *gorm.DB, app *newrelic.Application) {
	userHandler := user.UserInit(tx)
	router.Use(nrecho.Middleware(app))
	router.POST("/login", userHandler.Login)

	auth := router.Group("")
	auth.Use(middleware.JWTWithConfig(jwt.JWTConfig))
	auth.Use(middlewares.VerifyToken)
	router.POST("/register", userHandler.Register)

	auth.POST("/reset/password", userHandler.ResetPassword)

	auth.GET("/deneme/:id", func(c echo.Context) error {
		return nil
	})

}
