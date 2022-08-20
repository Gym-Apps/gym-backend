package user

import (
	"net/http"

	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Login(c echo.Context) error
}

type UserHandler struct {
	service userService.IUserService
}

func NewUserHandler(service userService.IUserService) IUserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login page2")
}
