package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"gorm.io/gorm"
)

var Handler IUserHandler

func UserInit(db *gorm.DB) {
	repository := userRepo.NewUserRepository(db)
	service := userService.NewUserService(repository)
	handler := NewUserHandler(service)

	Handler = handler
}
