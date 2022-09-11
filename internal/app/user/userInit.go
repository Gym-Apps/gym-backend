package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	userService "github.com/Gym-Apps/gym-backend/internal/service/user"
	"github.com/Gym-Apps/gym-backend/internal/utils"
	"gorm.io/gorm"
)

var Handler IUserHandler

func UserInit(db *gorm.DB) {
	utils := utils.NewUtils()
	repository := userRepo.NewUserRepository(db)
	service := userService.NewUserService(repository, utils)
	handler := NewUserHandler(service, utils)

	Handler = handler
}
