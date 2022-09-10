package user

import (
	"net/http"

	"github.com/Gym-Apps/gym-backend/dto/request"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/Gym-Apps/gym-backend/internal/utils"
	"github.com/Gym-Apps/gym-backend/internal/utils/response"
	"github.com/Gym-Apps/gym-backend/internal/utils/validate"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Login(c echo.Context) error
	ResetPassword(c echo.Context) error
}

type UserHandler struct {
	service userService.IUserService
	utils   utils.IUtils
}

func NewUserHandler(service userService.IUserService, utils utils.IUtils) IUserHandler {
	return &UserHandler{
		service: service,
		utils:   utils,
	}
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

func (h *UserHandler) ResetPassword(c echo.Context) error {
	var request request.UserResetPasswordDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}

	user := h.utils.GetUser(&c)

	err := h.service.ResetPassword(user, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusOK, response.Response(http.StatusOK, user))
}
