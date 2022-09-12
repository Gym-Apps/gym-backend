package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	"github.com/Gym-Apps/gym-backend/internal/service"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/Gym-Apps/gym-backend/internal/utils"
	"gorm.io/gorm"
)

func UserInit(db *gorm.DB) IUserHandler {
	utils := utils.NewUtils()
	repository := userRepo.NewUserRepository(db)
	service := userService.NewUserService(repository, service.Service{Utils: utils})
	handler := NewUserHandler(service, utils)

	return handler
}
