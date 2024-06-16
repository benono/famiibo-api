package repository

import (
	"go-rest-api/model"

	"gorm.io/gorm"
)

type ICategoryRepository interface {
	FindAll(categories *[]model.Category) error
	FindByID(category *model.Category, id uint) error
	Create(category *model.Category) error
	Update(category *model.Category) error
	Delete(id uint) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &categoryRepository{db}
}

func (cr *categoryRepository) FindAll(categories *[]model.Category) error {
	if err := cr.db.Find(categories).Error; err != nil {
		return err
	}
	return nil
}

func (cr *categoryRepository) FindByID(category *model.Category, id uint) error {
	if err := cr.db.First(category, id).Error; err != nil {
		return err
	}
	return nil
}

func (cr *categoryRepository) Create(category *model.Category) error {
	if err := cr.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (cr *categoryRepository) Update(category *model.Category) error {
	if err := cr.db.Save(category).Error; err != nil {
		return err
	}
	return nil
}

func (cr *categoryRepository) Delete(id uint) error {
	if err := cr.db.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
