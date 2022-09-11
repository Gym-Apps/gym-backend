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
	Register(userRegisterRequest request.UserRegisterDTO) (response.UserRegisterDTO, error)
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
func (s *UserService) Register(userRegisterRequest request.UserRegisterDTO) (response.UserRegisterDTO, error) {

	isDuplicateEmail := s.repository.IsDuplicateEmail(userRegisterRequest.Email)
	if isDuplicateEmail {
		err := errors.New("Bu e-mail adresi  farklı bir kullanıcı tarafından kullanılmaktadır.")
		return response.UserRegisterDTO{}, err
	}

	isDuplicatePhone := s.repository.IsDuplicatePhone(userRegisterRequest.Phone)
	if isDuplicatePhone {
		err := errors.New("Bu telefon numarası  farklı bir kullanıcı tarafından kullanılmaktadır.")
		return response.UserRegisterDTO{}, err
	}

	hashPassword, err := s.utils.GeneratePassword(userRegisterRequest.Password)
	if err != nil {
		err := errors.New("Şifre oluşturulamadı.")
		return response.UserRegisterDTO{}, err
	}

	passwordControl := s.utils.EqualPassword(hashPassword, userRegisterRequest.Password)
	if !passwordControl {
		return response.UserRegisterDTO{}, errors.New("Şifre doğru bir şekilde oluşturulamadı.")
	}

	birthday, err := utils.EuToTime(userRegisterRequest.Birthday)
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("Tarih dönüştürülürken hata oluştu")
	}

	newUser := models.User{
		Name:        userRegisterRequest.Name,
		Surname:     userRegisterRequest.Surname,
		Email:       userRegisterRequest.Email,
		Phone:       userRegisterRequest.Phone,
		Birthday:    birthday,
		AccountType: models.AccountType(userRegisterRequest.AccountType),
		AccountName: "Sporcu",
		Gender:      models.Gender(userRegisterRequest.Gender),
		GenderName:  "Erkek",
		Password:    string(hashPassword),
	}
	err = s.repository.Register(&newUser)
	if err != nil {
		err = errors.New("Kayıt işlemi başarısız oldu. Lütfen bilgilerinizi kontrol ediniz.")
		return response.UserRegisterDTO{}, err
	}
	var userRegisterResponse response.UserRegisterDTO
	userRegisterResponse.Convert(newUser)

	return userRegisterResponse, nil

}
