package user_management

import (
	"context"
	"github.com/andrepriyanto10/server_favaa/internal/model"
)

type UserContractRepository interface {
	StoreDataUser(ctx context.Context, user *model.User) error
	UpdateDataUser(ctx context.Context, email *string) error
	FetchUserByEmail(ctx context.Context, email string) (*model.User, error)
	//UpdatePassword(ctx context.Context, email string, pass string) error
}
