package user

import (
	userRepo "github.com/Gym-Apps/gym-backend/internal/repository/user"
)

type IUserService interface {
	Login() error
}

type UserService struct {
	repository userRepo.IUserRepository
}

func NewUserService(repository userRepo.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (s *UserService) Login() error {
	return nil
}
