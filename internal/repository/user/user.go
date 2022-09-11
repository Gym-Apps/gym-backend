package user

import (
	"sync"

	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login(phone string) (models.User, error)
	UpdatePassword(userID uint, password string) error
}

type UserRepository struct {
	mu sync.Mutex
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Login(phone string) (models.User, error) {
	u.mu.Lock()
	var user models.User
	err := u.db.Where("phone = ?", phone).First(&user).Error
	u.mu.Unlock()
	return user, err
}

func (u *UserRepository) UpdatePassword(userID uint, password string) error {
	u.mu.Lock()
	err := u.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
	u.mu.Unlock()
	return err
}
