package usecase

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IAccountUsecase interface {
	FindAll() ([]model.Account, error)
	FindById(id uint) (model.Account, error)
	Create(account model.Account) (model.Account, error)
	Update(account model.Account) (model.Account, error)
	Delete(id uint) error
}

type accountUsecase struct {
	cr repository.IAccountRepository
	cv validator.IAccountValidator
}

func NewAccountUsecase(cr repository.IAccountRepository, cv validator.IAccountValidator) IAccountUsecase {
	return &accountUsecase{cr, cv}
}

func (au *accountUsecase) FindAll() ([]model.Account, error) {
	fmt.Println("Executing FindAll method in accountUsecase")

	accounts := []model.Account{}
	if err := au.cr.FindAll(&accounts); err != nil {
		return []model.Account{}, err
	}
	return accounts, nil
}

func (au *accountUsecase) FindById(id uint) (model.Account, error) {
	account := model.Account{}
	if err := au.cr.FindById(&account, id); err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (au *accountUsecase) Create(account model.Account) (model.Account, error) {
	if err := au.cv.AccountValidate(account); err != nil {
		return model.Account{}, err
	}
	if err := au.cr.Create(&account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (au *accountUsecase) Update(account model.Account) (model.Account, error) {
	if err := au.cv.AccountValidate(account); err != nil {
		return model.Account{}, err
	}
	if err := au.cr.Update(&account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}

func (au *accountUsecase) Delete(id uint) error {
	if err := au.cr.Delete(id); err != nil {
		return err
	}
	return nil
}
