package utils

import "github.com/go-playground/validator/v10"

type ErrResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

func NewValidate(req any) []*ErrResponse {
	var errors []*ErrResponse
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, &ErrResponse{
				Field: err.StructNamespace(),
				Tag:   err.Tag(),
				Value: err.Param(),
			})
		}
	}

	return errors

}
