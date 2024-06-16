package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IPayeeRepository interface {
	FindAll(payees *[]model.Payee) error
	FindById(payee *model.Payee, id uint) error
	Create(payee *model.Payee) error
	Update(payee *model.Payee) error
	Delete(id uint) error
}

type payeeRepository struct {
	db *gorm.DB
}

func NewPayeeRepository(db *gorm.DB) IPayeeRepository {
	return &payeeRepository{db}
}

func (ur *payeeRepository) FindById(payee *model.Payee, id uint) error {
	if err := ur.db.Where("id=?", id).First(payee).Error; err != nil {
		return err
	}
	return nil
}

func (ur *payeeRepository) FindAll(payees *[]model.Payee) error {
	if err := ur.db.Find(payees).Error; err != nil {
		return err
	}
	return nil
}

func (ur *payeeRepository) Create(payee *model.Payee) error {
	if err := ur.db.Create(payee).Error; err != nil {
		return err
	}
	return nil
}

func (cr *payeeRepository) Update(payee *model.Payee) error {
	if err := cr.db.Save(payee).Error; err != nil {
		return err
	}
	return nil
}

func (cr *payeeRepository) Delete(id uint) error {
	if err := cr.db.Delete(&model.Payee{}, id).Error; err != nil {
		return err
	}
	return nil
}
