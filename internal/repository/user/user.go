package user

import (
	"github.com/Gym-Apps/gym-backend/internal/config/db"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login() error
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() IUserRepository {
	return &UserRepository{db: db.DB}
}

func (u UserRepository) Login() error {
	return nil
}
