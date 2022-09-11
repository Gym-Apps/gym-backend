package user

import (
	"context"
	"sync"

	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login(ctx context.Context, phone string) (models.User, error)
	UpdatePassword(ctx context.Context, userID uint, password string) error
	//WithContext(ctx context.Context) IUserRepository
}

type UserRepository struct {
	mu sync.Mutex
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

// func (u *UserRepository) WithContext(ctx context.Context) IUserRepository {
// 	u.db = u.db.WithContext(ctx)
// 	return u
// }

func (u *UserRepository) Login(ctx context.Context, phone string) (models.User, error) {
	u.mu.Lock()
	var user models.User
	err := u.db.Where("phone = ?", phone).First(&user).Error
	u.mu.Unlock()
	return user, err
}

func (u *UserRepository) UpdatePassword(ctx context.Context, userID uint, password string) error {
	u.mu.Lock()
	err := u.db.Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
	u.mu.Unlock()
	return err
}
