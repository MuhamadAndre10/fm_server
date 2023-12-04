package service

import (
	"context"
	"github.com/andrepriyanto10/server_favaa/configs/logger"
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/utils"
	"github.com/pkg/errors"
	"time"
)

type UserService struct {
	userRepo user_management.UserContractRepository
	log      *logger.Log
	timeout  time.Duration
}

func NewUserService(userRepo user_management.UserContractRepository, log *logger.Log) *UserService {
	return &UserService{
		userRepo: userRepo,
		log:      log,
		timeout:  time.Duration(4) * time.Second,
	}
}

func (s *UserService) Register(req *user_management.UserRegisterRequest) (*user_management.UserRegisterResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.Wrap(err, "failed to hash password")
	}

}
