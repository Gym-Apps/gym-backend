package user

import (
	"errors"
	"time"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/dto/response"
	jwtConfig "github.com/Gym-Apps/gym-backend/internal/config/jwt"
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	"github.com/Gym-Apps/gym-backend/internal/utils"
	jwtPackage "github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error)
	ResetPassword(user models.User, request request.UserResetPasswordDTO) error
}

type UserService struct {
	repository userRepo.IUserRepository
	utils      utils.IUtils
}

func NewUserService(repository userRepo.IUserRepository, utils utils.IUtils) IUserService {
	return &UserService{repository: repository, utils: utils}
}

func (s *UserService) Login(userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error) {
	user, err := s.repository.Login(userLoginRequest.Phone)
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

func (s *UserService) ResetPassword(user models.User, request request.UserResetPasswordDTO) error {
	// passwordControl := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
	passwordControl := s.utils.EqualPassword(user.Password, request.OldPassword)
	if !passwordControl {
		return errors.New("Eski şifre doğrulanamadı.")
	}

	equalPassword := s.utils.EqualPassword(user.Password, request.NewPassword)
	if equalPassword {
		return errors.New("Eski şifre ile yeni şifre aynı olamaz.")
	}

	password, _ := s.utils.GeneratePassword(request.NewPassword)

	err := s.repository.UpdatePassword(user.ID, string(password))
	if err != nil {
		return errors.New("şifre güncellemede sorun oluştu.")
	}

	return nil
}
