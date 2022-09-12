package user

import (
	"context"
	"sync"

	"github.com/Gym-Apps/gym-backend/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Login(ctx context.Context, phone string) (models.User, error)
	Register(ctx context.Context, user *models.User) error
	IsDuplicatePhone(ctx context.Context, phone string) bool
	IsDuplicateEmail(ctx context.Context, email string) bool
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

func (u *UserRepository) Login(ctx context.Context, phone string) (models.User, error) {
	u.mu.Lock()
	var user models.User
	err := u.db.WithContext(ctx).Where("phone = ?", phone).First(&user).Error
	u.mu.Unlock()
	return user, err
}

func (u *UserRepository) Register(ctx context.Context, user *models.User) error {
	u.mu.Lock()
	err := u.db.WithContext(ctx).Create(user).Error
	u.mu.Unlock()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) IsDuplicatePhone(ctx context.Context, phone string) bool {
	user := models.User{}
	u.mu.Lock()
	err := u.db.WithContext(ctx).Where("phone=?", phone).First(&user).Error
	u.mu.Unlock()
	if err != nil || user.ID <= 0 {
		return false
	}
	return true

}

func (u *UserRepository) IsDuplicateEmail(ctx context.Context, email string) bool {
	var user models.User
	u.mu.Lock()
	err := u.db.WithContext(ctx).Where("email=?", email).First(&user).Error
	u.mu.Unlock()
	if err != nil || user.ID <= 0 {
		return false
	}
	return true
}

func (u *UserRepository) UpdatePassword(ctx context.Context, userID uint, password string) error {
	u.mu.Lock()
	err := u.db.WithContext(ctx).Model(&models.User{}).Where("id = ?", userID).Update("password", password).Error
	u.mu.Unlock()
	return err
}
