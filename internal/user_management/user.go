package user_management

type UserRegisterRequest struct {
	FirstName string `json:"first_name,omitempty" validate:"required,alpha,min=3"`
	LastName  string `json:"last_name,omitempty" validate:"required,alpha,min=3"`
	Email     string `json:"email,omitempty" validate:"required,email"`
	Password  string `json:"password,omitempty" validate:"required,min=8,max=32"`
}

type UserRegisterResponse struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CodeRequest struct {
	Code string `json:"code,omitempty" validate:"required"`
}
