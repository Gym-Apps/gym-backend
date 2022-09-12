package user

import (
	"context"
	"errors"
	"testing"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/dto/response"
	"github.com/Gym-Apps/gym-backend/internal/utils"
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
		//repoMock.On("WithContext", mock.Anything).Return(&UserService{})
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
		//repoMock.On("WithContext", mock.Anything).Return(*repoMock)
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

func TestRegister(t *testing.T) {

	userRegisterRequest := request.UserRegisterDTO{
		Name:        "Test",
		Surname:     "test",
		Phone:       "5555555555",
		Email:       "example@example.com",
		Password:    "123456",
		Birthday:    "02.01.2006",
		AccountType: 2,
		Gender:      2,
	}

	t.Run("successful register", func(t *testing.T) {

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("GeneratePassword", userRegisterRequest.Password).Return("$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu", nil)
		utilsMock.On("EqualPassword", "$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu", userRegisterRequest.Password).Return(true)

		birthday, _ := utils.EuToTime(userRegisterRequest.Birthday)

		user := models.User{
			Name:        userRegisterRequest.Name,
			Surname:     userRegisterRequest.Surname,
			Phone:       userRegisterRequest.Phone,
			Email:       userRegisterRequest.Email,
			AccountType: models.AccountType(userRegisterRequest.AccountType),
			Gender:      models.Gender(userRegisterRequest.Gender),
			Birthday:    birthday,
			GenderName:  "Erkek",
			AccountName: "Sporcu",
			Password:    "$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu",
		}

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(false)
		repoMock.On("IsDuplicatePhone", userRegisterRequest.Phone).Return(false)
		repoMock.On("Register", &user).Return(nil)

		userService := NewUserService(repoMock, utilsMock)
		ctx := context.Background()
		userRegisterResponse, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, err, nil)
		assert.Equal(t, userRegisterResponse.Name, user.Name)

	})

	t.Run("duplicate e-mail", func(t *testing.T) {

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(true)

		userService := NewUserService(repoMock, nil)
		ctx := context.Background()

		userRegisterResponse, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, userRegisterResponse, response.UserRegisterDTO{})
		assert.Equal(t, err, errors.New("Bu e-mail adresi  farklı bir kullanıcı tarafından kullanılmaktadır."))
	})

	t.Run("duplicate phone", func(t *testing.T) {

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(false)
		repoMock.On("IsDuplicatePhone", userRegisterRequest.Phone).Return(true)

		userService := NewUserService(repoMock, nil)
		ctx := context.Background()
		userRegisterResponse, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, userRegisterResponse, response.UserRegisterDTO{})
		assert.Equal(t, err, errors.New("Bu telefon numarası  farklı bir kullanıcı tarafından kullanılmaktadır."))
	})

	t.Run("hash failed", func(t *testing.T) {

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("GeneratePassword", userRegisterRequest.Password).Return("", errors.New("Şifre oluşturulamadı."))

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(false)
		repoMock.On("IsDuplicatePhone", userRegisterRequest.Phone).Return(false)

		userService := NewUserService(repoMock, utilsMock)
		ctx := context.Background()
		userRegisterResponse, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, err, errors.New("Şifre oluşturulamadı."))
		assert.Equal(t, userRegisterResponse, response.UserRegisterDTO{})

	})
	t.Run("not equal", func(t *testing.T) {

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("GeneratePassword", userRegisterRequest.Password).Return("$2a$04$/gKgvJiNWsTt8VmXkb/Al.N5eOZDTaVTFlrjb4lgILIoguAyPj5Yq", nil)
		utilsMock.On("EqualPassword", "$2a$04$/gKgvJiNWsTt8VmXkb/Al.N5eOZDTaVTFlrjb4lgILIoguAyPj5Yq", userRegisterRequest.Password).Return(false)

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(false)
		repoMock.On("IsDuplicatePhone", userRegisterRequest.Phone).Return(false)

		userService := NewUserService(repoMock, utilsMock)
		ctx := context.Background()
		userRegisterResponse, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, err, errors.New("Şifre doğru bir şekilde oluşturulamadı."))
		assert.Equal(t, userRegisterResponse, response.UserRegisterDTO{})

	})
	t.Run("failed register", func(t *testing.T) {

		utilsMock := mocks.NewIUtils(t)
		utilsMock.On("GeneratePassword", userRegisterRequest.Password).Return("$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu", nil)
		utilsMock.On("EqualPassword", "$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu", userRegisterRequest.Password).Return(true)

		birthday, _ := utils.EuToTime(userRegisterRequest.Birthday)

		user := models.User{
			Name:        userRegisterRequest.Name,
			Surname:     userRegisterRequest.Surname,
			Phone:       userRegisterRequest.Phone,
			Email:       userRegisterRequest.Email,
			AccountType: models.AccountType(userRegisterRequest.AccountType),
			Gender:      models.Gender(userRegisterRequest.Gender),
			Birthday:    birthday,
			GenderName:  "Erkek",
			AccountName: "Sporcu",
			Password:    "$2a$04$tYJMolWtAagNYENlDvj/he5rvEkMsVYl6sKPvbA.W9xeCUoqKTDRu",
		}

		repoMock := mocks.NewIUserRepository(t)
		repoMock.On("IsDuplicateEmail", userRegisterRequest.Email).Return(false)
		repoMock.On("IsDuplicatePhone", userRegisterRequest.Phone).Return(false)
		repoMock.On("Register", &user).Return(errors.New("Kayıt işlemi başarısız oldu. Lütfen bilgilerinizi kontrol ediniz."))

		userService := NewUserService(repoMock, utilsMock)
		ctx := context.Background()
		_, err := userService.Register(ctx,userRegisterRequest)
		assert.Equal(t, err, errors.New("Kayıt işlemi başarısız oldu. Lütfen bilgilerinizi kontrol ediniz."))

	})

}
