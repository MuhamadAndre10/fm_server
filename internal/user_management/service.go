package user_management

// UserContractService is an interface that defines the contract of the user service
type UserContractService interface {
	Register(req *UserRegisterRequest) (*UserRegisterResponse, error)
	Verify(req *UserVerifyRequest) error
}

// MailService is an interface that defines the contract of the mail service
type MailService interface {
}
