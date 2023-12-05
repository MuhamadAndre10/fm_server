package service

import (
	"context"
	"encoding/json"
	"github.com/allegro/bigcache/v3"
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

func (s *UserService) Register(req *user_management.UserRegisterRequest, code *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	req.Password = hashPassword

	// save data in cache
	bigCache, err := bigcache.New(ctx, bigcache.DefaultConfig(5*time.Minute))
	if err != nil {
		return err
	}

	data := struct {
		FullName string
		Email    string
		Password string
		Code     *string
	}{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
		Code:     code,
	}

	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = cache.NewDataCache(bigCache).Set("user", dataByte)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) VerifyUserRegister(code *user_management.CodeRequest) error {
	data := cache.New()

	get, err := data.Get("user")
	if err != nil {
		return errors.WithMessage(err, "failed to get data from cache")
	}

	var user user_management.UserRegisterRequest
	err = json.Unmarshal(get, &user)
	if err != nil {
		return err
	}

	panic("implement me")

}
