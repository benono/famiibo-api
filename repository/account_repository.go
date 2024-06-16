package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type IAccountRepository interface {
	FindAll(accounts *[]model.Account) error
	FindById(account *model.Account, id uint) error
	Create(account *model.Account) error
	Update(account *model.Account) error
	Delete(id uint) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{db}
}

func (ur *accountRepository) FindById(account *model.Account, id uint) error {
	if err := ur.db.Where("id=?", id).First(account).Error; err != nil {
		return err
	}
	return nil
}

func (ur *accountRepository) FindAll(accounts *[]model.Account) error {
	if err := ur.db.Find(accounts).Error; err != nil {
		return err
	}
	return nil
}

func (ur *accountRepository) Create(account *model.Account) error {
	if err := ur.db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (cr *accountRepository) Update(account *model.Account) error {
	if err := cr.db.Save(account).Error; err != nil {
		return err
	}
	return nil
}

func (cr *accountRepository) Delete(id uint) error {
	if err := cr.db.Delete(&model.Account{}, id).Error; err != nil {
		return err
	}
	return nil
}
