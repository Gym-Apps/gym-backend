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
	"github.com/Gym-Apps/gym-backend/internal/utils"
	jwtPackage "github.com/Gym-Apps/gym-backend/internal/utils/jwt"
	"github.com/Gym-Apps/gym-backend/models"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Login(ctx context.Context, userLoginRequest request.UserLoginDTO) (response.UserLoginDTO, error)
	Register(ctx context.Context, userRegisterRequest request.UserRegisterDTO) (response.UserRegisterDTO, error)
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
		return errors.New("Eski ??ifre do??rulanamad??.")
	}

	equalPassword := s.Service.Utils.EqualPassword(user.Password, request.NewPassword)
	if equalPassword {
		return errors.New("Eski ??ifre ile yeni ??ifre ayn?? olamaz.")
	}

	password, _ := s.Service.Utils.GeneratePassword(request.NewPassword)

	err := s.repository.UpdatePassword(ctx, user.ID, string(password))
	if err != nil {
		return errors.New("??ifre g??ncellemede sorun olu??tu.")
	}

	return nil
}
func (s *UserService) Register(ctx context.Context, userRegisterRequest request.UserRegisterDTO) (response.UserRegisterDTO, error) {

	ctx, cancel := context.WithTimeout(ctx, db.Time)
	defer cancel()

	isDuplicateEmail := s.repository.IsDuplicateEmail(ctx,userRegisterRequest.Email)
	if isDuplicateEmail {
		err := errors.New("Bu e-mail adresi  farkl?? bir kullan??c?? taraf??ndan kullan??lmaktad??r.")
		return response.UserRegisterDTO{}, err
	}

	isDuplicatePhone := s.repository.IsDuplicatePhone(ctx,userRegisterRequest.Phone)
	if isDuplicatePhone {
		err := errors.New("Bu telefon numaras??  farkl?? bir kullan??c?? taraf??ndan kullan??lmaktad??r.")
		return response.UserRegisterDTO{}, err
	}

	hashPassword, err := s.Utils.GeneratePassword(userRegisterRequest.Password)
	if err != nil {
		err := errors.New("??ifre olu??turulamad??.")
		return response.UserRegisterDTO{}, err
	}

	passwordControl := s.Utils.EqualPassword(hashPassword, userRegisterRequest.Password)
	if !passwordControl {
		return response.UserRegisterDTO{}, errors.New("??ifre do??ru bir ??ekilde olu??turulamad??.")
	}

	birthday, err := utils.EuToTime(userRegisterRequest.Birthday)
	if err != nil {
		return response.UserRegisterDTO{}, errors.New("Tarih d??n????t??r??l??rken hata olu??tu")
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
	err = s.repository.Register(ctx,&newUser)
	if err != nil {
		err = errors.New("Kay??t i??lemi ba??ar??s??z oldu. L??tfen bilgilerinizi kontrol ediniz.")
		return response.UserRegisterDTO{}, err
	}
	var userRegisterResponse response.UserRegisterDTO
	userRegisterResponse.Convert(newUser)

	return userRegisterResponse, nil

}
