package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IAccountValidator interface {
	AccountValidate(account model.Account) error
}

type accountValidator struct{}

func NewAccountValidator() IAccountValidator {
	return &accountValidator{}
}

func (cv *accountValidator) AccountValidate(account model.Account) error {
	return validation.ValidateStruct(&account,
		validation.Field(
			&account.Name,
			validation.Required.Error("name is required"),
		),
	)
}
