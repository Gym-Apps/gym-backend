package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
)

var Handler IUserHandler

func init() {
	repository := userRepo.NewUserRepository()
	service := userService.NewUserService(repository)
	handler := NewUserHandler(service)

	Handler = handler
}
