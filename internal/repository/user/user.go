package user

import (
	"sync"

	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login(phone string) (models.User, error)
	Register(user *models.User) error
	IsDuplicatePhone(phone string) bool
    IsDuplicateEmail(email string) bool 
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

func (u *UserRepository) Register(user *models.User) error {
	u.mu.Lock()
	err := u.db.Create(user).Error
	u.mu.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) IsDuplicatePhone(phone string) bool {
	user := models.User{}
	err := u.db.Where("phone=?", phone).First(&user).Error
	if err != nil || user.ID <= 0 {
		return false
	}
	return true

}

func (u *UserRepository) IsDuplicateEmail(email string) bool {
	var user models.User
	err := u.db.Where("email=?", email).First(&user).Error
	if err != nil || user.ID <= 0 {
		return false
	}
	return true
}

func (u *UserRepository) UpdatePassword(userID uint, password string) error {
	u.mu.Lock()
	err := u.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
	u.mu.Unlock()
	return err
}
