package service

import (
	"context"
	"encoding/json"
	"github.com/andrepriyanto10/server_favaa/internal/model"
	"github.com/andrepriyanto10/server_favaa/internal/user_management"
	"github.com/andrepriyanto10/server_favaa/pkg/cache"
	"github.com/andrepriyanto10/server_favaa/utils"
	"github.com/pkg/errors"
	"time"
)

type UserService struct {
	userRepo user_management.UserContractRepository
	timeout  time.Duration
}

func NewUserService(userRepo user_management.UserContractRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
		timeout:  time.Duration(4) * time.Second,
	}
}

func (s *UserService) Register(ctx context.Context, req *user_management.UserRegisterRequest, code *string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	req.Password = hashPassword

	user := &model.User{
		Password: req.Password,
		Email:    req.Email,
		MitraIdentity: &model.MitraIdentity{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	}

	err = s.userRepo.StoreDataUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) VerifyUserRegister(ctx context.Context, code *user_management.CodeRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	getDataCache, err := cache.Cache.Get("user")
	if err != nil {
		return errors.WithMessage(err, "failed to get data from dataCache")
	}

	var data struct {
		Email     string
		Code      string
		ExpiredAt time.Time
	}

	err = json.Unmarshal(getDataCache, &data)
	if err != nil {
		return err
	}

	if code.Code != data.Code {
		return errors.New("code not match")
	}

	if time.Now().After(data.ExpiredAt) {
		return errors.New("code expired")
	}

	err = s.userRepo.UpdateDataUser(ctx, &data.Email)
	if err != nil {
		return errors.WithMessage(err, "failed to store data user")
	}

	return nil

}
