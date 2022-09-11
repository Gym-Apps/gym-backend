package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/Gym-Apps/gym-backend/internal/utils"
	"gorm.io/gorm"
)

func UserInit(db *gorm.DB) IUserHandler {
	utils := utils.NewUtils()
	repository := userRepo.NewUserRepository(db)
	service := userService.NewUserService(repository, utils)
	handler := NewUserHandler(service, utils)

	return handler
}
