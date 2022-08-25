package user

import (
	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login(phone string) (models.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Login(phone string) (models.User, error) {
	var user models.User
	err := u.db.Where("phone = ?", phone).First(&user).Error
	return user, err
}
