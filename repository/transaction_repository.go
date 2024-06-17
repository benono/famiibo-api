package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ITransactionRepository interface {
	FindAll(transactions *[]model.TransactionResponse) error
	FindById(transaction *model.TransactionResponse, id uint) error
	Create(transaction *model.Transaction) error
	Update(transaction *model.Transaction) error
	Delete(id uint) error
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) ITransactionRepository {
	return &transactionRepository{db}
}

func (ur *transactionRepository) FindById(transaction *model.TransactionResponse, id uint) error {
	query := ur.db.Model(&model.Transaction{})
	query = query.Joins("JOIN categories ON transactions.category_id = categories.id")
	query = query.Joins("JOIN accounts ON transactions.account_id = accounts.id")
	query = query.Joins("JOIN payees ON transactions.payee_id = payees.id")
	query = query.Where("transactions.id=?", id)
	query = query.First(transaction)
	if err := query.Error; err != nil {
		return err
	}
	return nil
}

func (ur *transactionRepository) FindAll(transactions *[]model.TransactionResponse) error {
	query := ur.db.Model(&model.Transaction{})
	query = query.Joins("JOIN categories ON transactions.category_id = categories.id")
	query = query.Joins("JOIN accounts ON transactions.account_id = accounts.id")
	query = query.Joins("JOIN payees ON transactions.payee_id = payees.id")
	if err := query.Find(transactions).Error; err != nil {
		return err
	}
	return nil
}

func (ur *transactionRepository) Create(transaction *model.Transaction) error {
	if err := ur.db.Create(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (cr *transactionRepository) Update(transaction *model.Transaction) error {
	if err := cr.db.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}

func (cr *transactionRepository) Delete(id uint) error {
	if err := cr.db.Delete(&model.Transaction{}, id).Error; err != nil {
		return err
	}
	return nil
}
