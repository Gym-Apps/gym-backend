package user

import (
	"net/http"

	"github.com/Gym-Apps/gym-backend/dto/request"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/Gym-Apps/gym-backend/internal/util/response"
	"github.com/Gym-Apps/gym-backend/internal/util/validate"

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
	var request request.UserLoginDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}

	user, err := h.service.Login(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, "Bir hata oluştu"))
	}
	return c.JSON(http.StatusOK, response.Response(http.StatusOK, user))
}
