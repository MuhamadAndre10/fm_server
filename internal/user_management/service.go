package user_management

import "context"

// UserContractService is an interface that defines the contract of the user service
type UserContractService interface {
	Register(ctx context.Context, req *UserRegisterRequest, code *string) error
	VerifyUserRegister(ctx context.Context, code *CodeRequest) error
}

// MailService is an interface that defines the contract of the mail service
type MailService interface {
	SendMailWithSmtp(to []string, subject, body string) error
	SendMailWithSendGrid(to []string, subject, body string) error
}
