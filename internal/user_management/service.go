package user_management

// UserContractService is an interface that defines the contract of the user service
type UserContractService interface {
	Register(req *UserRegisterRequest) error
	Verify(code *CodeRequest) error
}

// MailService is an interface that defines the contract of the mail service
type MailService interface {
	SendMailWithSmtp(to []string, subject, body string) error
	SendMailWithSendGrid(to []string, subject, body string) error
}
