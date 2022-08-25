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
		handler := NewUserHandler(serviceMock)

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
		handler := NewUserHandler(serviceMock)

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
		handler := NewUserHandler(serviceMock)

		// Assertions
		if assert.NoError(t, handler.Login(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
