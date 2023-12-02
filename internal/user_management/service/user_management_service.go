package service

import "github.com/andrepriyanto10/server_favaa/internal/user_management"

type UserService struct {
	userRepo user_management.UserContractRepository
}

func NewUserService(userRepo user_management.UserContractRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}
