package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPayeeValidator interface {
	PayeeValidate(payee model.Payee) error
}

type payeeValidator struct{}

func NewPayeeValidator() IPayeeValidator {
	return &payeeValidator{}
}

func (uv *payeeValidator) PayeeValidate(payee model.Payee) error {
	return validation.ValidateStruct(&payee,
		validation.Field(
			&payee.Name,
			validation.Required.Error("name is required"),
		),
	)
}
