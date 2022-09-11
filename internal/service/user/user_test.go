package user

import (
	"context"
	"errors"
	"testing"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/mocks"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	var userLoginRequest request.UserLoginDTO
	userLoginRequest.Phone = "5551755445"
	userLoginRequest.Password = "123123"

	t.Run("first test", func(t *testing.T) {
		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("Login", mock.Anything, userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO",
		}, nil)

		userService := NewUserService(repoMock, nil)

		userLoginResponse, err := userService.Login(ctx, userLoginRequest)
		assert.Equal(t, err, nil)
		assert.Equal(t, userLoginResponse.Phone, userLoginRequest.Phone)
	})

	t.Run("second test", func(t *testing.T) {
		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("Login", mock.Anything, userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO",
		}, errors.New("not found"))

		userService := NewUserService(repoMock, nil)

		_, err := userService.Login(ctx, userLoginRequest)
		assert.NotEqual(t, err, nil)
	})

	t.Run("third test", func(t *testing.T) {
		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)
		//repoMock.On("WithContext", mock.Anything).Return(*repoMock)
		repoMock.On("Login", mock.Anything, userLoginRequest.Phone).Return(models.User{
			Name:       "baran",
			Surname:    "atbaş",
			Gender:     1,
			GenderName: "Erkek",
			Phone:      userLoginRequest.Phone,
			Password:   "$2a$04$HnXu0HWPzJlFR6R5g2Nn",
		}, nil)

		userService := NewUserService(repoMock, nil)

		_, err := userService.Login(ctx, userLoginRequest)
		assert.NotEqual(t, err, nil)
	})
}

func TestResetPassword(t *testing.T) {
	t.Run("Successful reset password", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		user.ID = 12

		var request request.UserResetPasswordDTO
		request.OldPassword = "123123"
		request.NewPassword = "123456"

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("EqualPassword", user.Password, request.OldPassword).Return(true)
		utilsMock.On("EqualPassword", user.Password, request.NewPassword).Return(false)
		utilsMock.On("GeneratePassword", request.NewPassword).Return("$04$1kp13XtORrd5gI0Buf.3ceUN/Ee94Ok0L.1AMwJBEAHoBZFRNPo7S", nil)

		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("UpdatePassword", mock.Anything, user.ID, "$04$1kp13XtORrd5gI0Buf.3ceUN/Ee94Ok0L.1AMwJBEAHoBZFRNPo7S").Return(nil)

		userService := NewUserService(repoMock, utilsMock)
		err := userService.ResetPassword(ctx, user, request)
		assert.Equal(t, err, nil)
	})

	t.Run("Worng old password", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		user.ID = 12

		var request request.UserResetPasswordDTO
		request.OldPassword = "1231232"
		request.NewPassword = "123456"

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("EqualPassword", user.Password, request.OldPassword).Return(false)

		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)

		userService := NewUserService(repoMock, utilsMock)
		err := userService.ResetPassword(ctx, user, request)
		assert.Equal(t, err, errors.New("Eski şifre doğrulanamadı."))
	})

	t.Run("Equal old password and new password", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		user.ID = 12

		var request request.UserResetPasswordDTO
		request.OldPassword = "123123"
		request.NewPassword = "123123"

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("EqualPassword", user.Password, request.OldPassword).Return(true)
		utilsMock.On("EqualPassword", user.Password, request.NewPassword).Return(true)

		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)

		userService := NewUserService(repoMock, utilsMock)
		err := userService.ResetPassword(ctx, user, request)
		assert.Equal(t, err, errors.New("Eski şifre ile yeni şifre aynı olamaz."))
	})

	t.Run("Repository error", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		user.ID = 12

		var request request.UserResetPasswordDTO
		request.OldPassword = "123123"
		request.NewPassword = "123456"

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("EqualPassword", user.Password, request.OldPassword).Return(true)
		utilsMock.On("EqualPassword", user.Password, request.NewPassword).Return(false)
		utilsMock.On("GeneratePassword", request.NewPassword).Return("$04$1kp13XtORrd5gI0Buf.3ceUN/Ee94Ok0L.1AMwJBEAHoBZFRNPo7S", nil)

		ctx := context.Background()
		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("UpdatePassword", mock.Anything, user.ID, "$04$1kp13XtORrd5gI0Buf.3ceUN/Ee94Ok0L.1AMwJBEAHoBZFRNPo7S").Return(errors.New("şifre güncellemede sorun oluştu."))

		userService := NewUserService(repoMock, utilsMock)
		err := userService.ResetPassword(ctx, user, request)
		assert.Equal(t, err, errors.New("şifre güncellemede sorun oluştu."))
	})
}
