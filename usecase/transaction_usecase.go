package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ITransactionUsecase interface {
	FindAll() ([]model.TransactionResponse, error)
	FindById(id uint) (model.TransactionResponse, error)
	Create(transaction model.Transaction) (model.Transaction, error)
	Update(transaction model.Transaction) (model.Transaction, error)
	Delete(id uint) error
}

type transactionUsecase struct {
	cr repository.ITransactionRepository
	cv validator.ITransactionValidator
}

func NewTransactionUsecase(cr repository.ITransactionRepository, cv validator.ITransactionValidator) ITransactionUsecase {
	return &transactionUsecase{cr, cv}
}

func (pu *transactionUsecase) FindAll() ([]model.TransactionResponse, error) {
	transactions := []model.TransactionResponse{}
	if err := pu.cr.FindAll(&transactions); err != nil {
		return []model.TransactionResponse{}, err
	}
	return transactions, nil
}

func (pu *transactionUsecase) FindById(id uint) (model.TransactionResponse, error) {
	transaction := model.TransactionResponse{}
	if err := pu.cr.FindById(&transaction, id); err != nil {
		return model.TransactionResponse{}, err
	}
	return transaction, nil
}

func (pu *transactionUsecase) Create(transaction model.Transaction) (model.Transaction, error) {
	if err := pu.cv.TransactionValidate(transaction); err != nil {
		return model.Transaction{}, err
	}
	if err := pu.cr.Create(&transaction); err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (pu *transactionUsecase) Update(transaction model.Transaction) (model.Transaction, error) {
	if err := pu.cv.TransactionValidate(transaction); err != nil {
		return model.Transaction{}, err
	}
	if err := pu.cr.Update(&transaction); err != nil {
		return model.Transaction{}, err
	}
	return transaction, nil
}

func (pu *transactionUsecase) Delete(id uint) error {
	if err := pu.cr.Delete(id); err != nil {
		return err
	}
	return nil
}
