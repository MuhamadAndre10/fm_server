package user_management

import "context"

// UserContractService is an interface that defines the contract of the user service
type UserContractService interface {
	Register(ctx context.Context, req *UserRegisterRequest, code *string) error
	VerifyUserRegister(ctx context.Context, code *CodeRequest) error
	Login(ctx context.Context, req *UserLoginRequest) error
	RecoveryPassword(ctx context.Context, email string) error
	//UpdatePassword(ctx context.Context, req *UpdatePassRequest) error
}

// MailService is an interface that defines the contract of the mail service
type MailService interface {
	SendMailWithSmtp(ctx context.Context, to []string, subject, body string) error
	SendMailWithSendGrid(ctx context.Context, to []string, subject, body string) error
}
