package user

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/dto/response"
	mocks "github.com/Gym-Apps/gym-backend/mocks"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setup(method string, path string, body io.Reader) {

}

func TestLogin(t *testing.T) {
	t.Run("first login", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)
		var loginRequest request.UserLoginDTO
		loginRequest.Password = "123123"
		loginRequest.Phone = "5551755445"

		serviceMock.On("Login", loginRequest).Return(response.UserLoginDTO{
			ID:      1,
			Name:    "baran",
			Surname: "atbaş",
			Phone:   "5551755445",
		}, nil)

		userJSON := `{"phone": "5551755445", "password": "123123"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		// Assertions
		if assert.NoError(t, handler.Login(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("second login", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)
		var loginRequest request.UserLoginDTO
		loginRequest.Phone = "5551755445"
		loginRequest.Password = "123123"

		userJSON := `{"phone": "5551755445"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		// Assertions
		if assert.NoError(t, handler.Login(c)) {
			assert.Equal(t, 400, rec.Code)
		}
		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})

	t.Run("third login", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)
		var loginRequest request.UserLoginDTO
		loginRequest.Password = "123123"
		loginRequest.Phone = "5551755445"

		serviceMock.On("Login", loginRequest).Return(response.UserLoginDTO{
			ID:      1,
			Name:    "baran",
			Surname: "atbaş",
			Phone:   "5551755445",
		}, errors.New("not found"))

		userJSON := `{"phone": "5551755445", "password": "123123"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		// Assertions
		if assert.NoError(t, handler.Login(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestResetPassword(t *testing.T) {

	t.Run("successful reset password", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		utilsMock := mocks.NewIUtils(t)
		serviceMock := mocks.NewIUserService(t)
		var resetPasswordRequest request.UserResetPasswordDTO
		resetPasswordRequest.OldPassword = "123456"
		resetPasswordRequest.NewPassword = "123123"

		serviceMock.On("ResetPassword", user, resetPasswordRequest).Return(nil)

		resetPasswordJSON := `{"old_password": "123456", "new_password": "123123"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/reset/password", strings.NewReader(resetPasswordJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		utilsMock.On("GetUser", &c).Return(user)

		handler := NewUserHandler(serviceMock, utilsMock)

		// Assertions
		if assert.NoError(t, handler.ResetPassword(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})

	t.Run("validate error", func(t *testing.T) {

		utilsMock := mocks.NewIUtils(t)
		serviceMock := mocks.NewIUserService(t)
		var resetPasswordRequest request.UserResetPasswordDTO
		resetPasswordRequest.OldPassword = "123456"
		resetPasswordRequest.NewPassword = "123123"

		//serviceMock.On("ResetPassword", user, resetPasswordRequest).Return(nil)

		resetPasswordJSON := `{"old_password": 12, "new_password": "123123"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/reset/password", strings.NewReader(resetPasswordJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		//utilsMock.On("GetUser", &c).Return(user)

		handler := NewUserHandler(serviceMock, utilsMock)

		// Assertions
		if assert.NoError(t, handler.ResetPassword(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("service error", func(t *testing.T) {
		user := models.User{
			Name:     "test",
			Surname:  "test surname",
			Password: "$2a$04$HnXu0HWPzJlFR6R5g2K81OywH.roBdJn1Ms7jiqua2yx38aI2zNnO", // 123123 demek
		}
		utilsMock := mocks.NewIUtils(t)
		serviceMock := mocks.NewIUserService(t)
		var resetPasswordRequest request.UserResetPasswordDTO
		resetPasswordRequest.OldPassword = "123456"
		resetPasswordRequest.NewPassword = "123123"

		serviceMock.On("ResetPassword", user, resetPasswordRequest).Return(errors.New("some error"))

		resetPasswordJSON := `{"old_password": "123456", "new_password": "123123"}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/reset/password", strings.NewReader(resetPasswordJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		utilsMock.On("GetUser", &c).Return(user)

		handler := NewUserHandler(serviceMock, utilsMock)

		// Assertions
		if assert.NoError(t, handler.ResetPassword(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}

func TestRegister(t *testing.T) {
	t.Run("successful register", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)

		request := request.UserRegisterDTO{
			Name:        "Test",
			Surname:     "test",
			Phone:       "5555555555",
			Email:       "example@example.com",
			Password:    "123456",
			Birthday:    "02.01.2006",
			AccountType: 2,
			Gender:      2,
		}

		serviceMock.On("Register", request).Return(response.UserRegisterDTO{
			Name:        request.Name,
			Surname:     request.Surname,
			Phone:       request.Phone,
			Email:       request.Email,
			Gender:      "Erkek",
			AccountName: "Sporcu",
		}, nil)

		userJSON := `
		{
		"name":"Test",
		"surname":"test",
		"phone":"5555555555",
		"password":"123456",
		"email":"example@example.com",
		"account_type" :2,
		"gender":2,
		"birthday":"02.01.2006" 
		}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		if assert.NoError(t, handler.Register(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)

		}

	})

	t.Run("validate error", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)

		userJSON := `{
		"name":"Test",
		"surname":"test",
		"phone":"5555555555",
		"password":"123456", }`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		if assert.NoError(t, handler.Register(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("service error", func(t *testing.T) {
		serviceMock := mocks.NewIUserService(t)

		request := request.UserRegisterDTO{
			Name:        "Test",
			Surname:     "test",
			Phone:       "5555555555",
			Email:       "example@example.com",
			Password:    "123456",
			Birthday:    "02.01.2006",
			AccountType: 2,
			Gender:      2,
		}

		serviceMock.On("Register", request).Return(response.UserRegisterDTO{}, errors.New("Service Error"))

		userJSON := `
		{
		"name":"Test",
		"surname":"test",
		"phone":"5555555555",
		"password":"123456",
		"email":"example@example.com",
		"account_type" :2,
		"gender":2,
		"birthday":"02.01.2006" 
		}`
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(userJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		handler := NewUserHandler(serviceMock, nil)

		if assert.NoError(t, handler.Register(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)

		}

	})

}
