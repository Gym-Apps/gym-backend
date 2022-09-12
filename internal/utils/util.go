package utils

import (
	"time"

	"github.com/Gym-Apps/gym-backend/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type IUtils interface {
	GetUser(c *echo.Context) models.User
	EqualPassword(old, new string) bool
	GeneratePassword(password string) (string, error)
	//EuToTime(StringDate string) (time.Time, error)
}

type Utils struct{}

func NewUtils() IUtils {
	return &Utils{}
}

func (u *Utils) GetUser(c *echo.Context) models.User {
	user := *(*c).Get("verifiedUser").(*models.User)
	return user
}

func (u *Utils) EqualPassword(old, new string) bool {
	passwordControl := bcrypt.CompareHashAndPassword([]byte(old), []byte(new))
	if passwordControl != nil {
		return false
	}
	return true
}

func (u *Utils) GeneratePassword(password string) (string, error) {
	hasPassword, err := bcrypt.GenerateFromPassword([]byte(password), 4)
	return string(hasPassword), err
}

func  EuToTime(StringDate string) (time.Time, error) {
	date, err := time.Parse("02.01.2006", StringDate)
	return date, err
}
