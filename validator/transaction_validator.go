package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITransactionValidator interface {
	TransactionValidate(transaction model.Transaction) error
}

type transactionValidator struct{}

func NewTransactionValidator() ITransactionValidator {
	return &transactionValidator{}
}

func (uv *transactionValidator) TransactionValidate(transaction model.Transaction) error {
	return validation.ValidateStruct(&transaction,
		validation.Field(
			&transaction.Amount,
			validation.Required.Error("amount is required"),
		),
		validation.Field(
			&transaction.Date,
			validation.Required.Error("date is required"),
		),
		validation.Field(
			&transaction.AccountID,
			validation.Required.Error("account_id is required"),
		),
	)
}
