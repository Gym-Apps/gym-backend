package user

import (
	"time"

	"github.com/Gym-Apps/gym-backend/dto/request"
	"github.com/Gym-Apps/gym-backend/dto/response"
	jwtConfig "github.com/Gym-Apps/gym-backend/internal/config/jwt"
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
	jwtPackage "github.com/Gym-Apps/gym-backend/internal/util/jwt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error)
}

type UserService struct {
	repository userRepo.IUserRepository
}

func NewUserService(repository userRepo.IUserRepository) IUserService {
	return &UserService{repository: repository}
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
