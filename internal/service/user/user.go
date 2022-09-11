package user

import (
	"context"
	"errors"
	"time"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/dto/response"
	"github.com/Gym-Apps/gym-backend/internal/config/db"
	jwtConfig "github.com/Gym-Apps/gym-backend/internal/config/jwt"
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	"github.com/Gym-Apps/gym-backend/internal/service"
	jwtPackage "github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(ctx context.Context, userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error)
	ResetPassword(ctx context.Context, user models.User, request request.UserResetPasswordDTO) error
}

type UserService struct {
	repository userRepo.IUserRepository
	service.Service
}

func NewUserService(repository userRepo.IUserRepository, service service.Service) IUserService {
	return &UserService{repository: repository, Service: service}
}

func (s *UserService) Login(ctx context.Context, userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Time)
	defer cancel()

	user, err := s.repository.Login(ctx, userLoginRequest.Phone)
	if err != nil {
		return response.UserLoginDTO{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return response.UserLoginDTO{}, err
	}

	claims := &jwtPackage.JwtCustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(jwtConfig.TokenSecret))
	if err != nil {
		return response.UserLoginDTO{}, err
	}

	var userLoginResponse response.UserLoginDTO
	userLoginResponse.Convert(user)
	userLoginResponse.Token = tokenSigned
	return userLoginResponse, nil
}

func (s *UserService) ResetPassword(ctx context.Context, user models.User, request request.UserResetPasswordDTO) error {
	ctx, cancel := context.WithTimeout(ctx, db.Time)
	defer cancel()
	passwordControl := s.Service.Utils.EqualPassword(user.Password, request.OldPassword)
	if !passwordControl {
		return errors.New("Eski şifre doğrulanamadı.")
	}

	equalPassword := s.Service.Utils.EqualPassword(user.Password, request.NewPassword)
	if equalPassword {
		return errors.New("Eski şifre ile yeni şifre aynı olamaz.")
	}

	password, _ := s.Service.Utils.GeneratePassword(request.NewPassword)

	err := s.repository.UpdatePassword(ctx, user.ID, string(password))
	if err != nil {
		return errors.New("şifre güncellemede sorun oluştu.")
	}

	return nil
}
