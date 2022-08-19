package main

import (
	"github.com/Gym-Apps/gym-backend/internal/config/db"
	"github.com/Gym-Apps/gym-backend/router"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()

	e := echo.New()
	router.Init(e)
	e.Logger.Fatal(e.Start(":8080"))
}
