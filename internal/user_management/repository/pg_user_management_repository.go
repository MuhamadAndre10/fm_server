package repository

import (
	"context"
	"github.com/andrepriyanto10/server_favaa/internal/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) StoreDataUser(ctx context.Context, user *model.User) error {

	usr := &model.User{
		Email:    user.Email,
		Password: user.Password,
	}

	err := u.DB.WithContext(ctx).Create(&usr).Error
	if err != nil {
		return err
	}

	err = u.DB.WithContext(ctx).Create(&model.MitraIdentity{
		UserID:    usr.ID,
		FirstName: user.MitraIdentity.FirstName,
		LastName:  user.MitraIdentity.LastName,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) UpdateDataUser(ctx context.Context, email *string) error {

	err := u.DB.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Update("verification_status", true).Error
	if err != nil {
		return err
	}

	return nil
}
