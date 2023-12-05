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

func (s *UserService) Register(req *user_management.UserRegisterRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), s.timeout)
	defer cancel()

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return errors.Wrap(err, "failed to hash password")
	}

	req.Password = hashPassword

	// save data in cache
	bigCache, err := bigcache.New(ctx, bigcache.DefaultConfig(10*time.Minute))
	if err != nil {
		return err
	}

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	err = cache.NewDataCache(bigCache).Set("user", data)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Verify(code *user_management.CodeRequest) error {
	panic("implement me")
}
