package user

import (
	"errors"
	"testing"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/mocks"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	var userLoginRequest request.UserLoginDTO
	userLoginRequest.Phone = "5551755445"
	userLoginRequest.Password = "123123"

	t.Run("first test", func(t *testing.T) {
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("Login", userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO",
		}, nil)

		userService := NewUserService(repoMock)

		userLoginResponse, err := userService.Login(userLoginRequest)
		assert.Equal(t, err, nil)
		assert.Equal(t, userLoginResponse.Phone, userLoginRequest.Phone)
	})

	t.Run("second test", func(t *testing.T) {
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("Login", userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO",
		}, errors.New("not found"))

		userService := NewUserService(repoMock)

		_, err := userService.Login(userLoginRequest)
		assert.NotEqual(t, err, nil)
	})

	t.Run("third test", func(t *testing.T) {
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("Login", userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2Nn",
		}, nil)

		userService := NewUserService(repoMock)

		_, err := userService.Login(userLoginRequest)
		assert.NotEqual(t, err, nil)
	})
}
