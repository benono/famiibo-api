package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type IPayeeUsecase interface {
	FindAll() ([]model.Payee, error)
	FindById(id uint) (model.Payee, error)
	Create(payee model.Payee) (model.Payee, error)
	Update(payee model.Payee) (model.Payee, error)
	Delete(id uint) error
}

type payeeUsecase struct {
	cr repository.IPayeeRepository
	cv validator.IPayeeValidator
}

func NewPayeeUsecase(cr repository.IPayeeRepository, cv validator.IPayeeValidator) IPayeeUsecase {
	return &payeeUsecase{cr, cv}
}

func (pu *payeeUsecase) FindAll() ([]model.Payee, error) {
	payees := []model.Payee{}
	if err := pu.cr.FindAll(&payees); err != nil {
		return []model.Payee{}, err
	}
	return payees, nil
}

func (pu *payeeUsecase) FindById(id uint) (model.Payee, error) {
	payee := model.Payee{}
	if err := pu.cr.FindById(&payee, id); err != nil {
		return model.Payee{}, err
	}
	return payee, nil
}

func (pu *payeeUsecase) Create(payee model.Payee) (model.Payee, error) {
	if err := pu.cv.PayeeValidate(payee); err != nil {
		return model.Payee{}, err
	}
	if err := pu.cr.Create(&payee); err != nil {
		return model.Payee{}, err
	}
	return payee, nil
}

func (pu *payeeUsecase) Update(payee model.Payee) (model.Payee, error) {
	if err := pu.cv.PayeeValidate(payee); err != nil {
		return model.Payee{}, err
	}
	if err := pu.cr.Update(&payee); err != nil {
		return model.Payee{}, err
	}
	return payee, nil
}

func (pu *payeeUsecase) Delete(id uint) error {
	if err := pu.cr.Delete(id); err != nil {
		return err
	}
	return nil
}
