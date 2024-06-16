package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
	"go-rest-api/validator"
)

type ICategoryUsecase interface {
	FindAll() ([]model.Category, error)
	FindByID(id uint) (model.Category, error)
	Create(category model.Category) (model.Category, error)
	Update(category model.Category) (model.Category, error)
	Delete(id uint) error
}

type categoryUsecase struct {
	cr repository.ICategoryRepository
	cv validator.ICategoryValidator
}

func NewCategoryUsecase(cr repository.ICategoryRepository, cv validator.ICategoryValidator) ICategoryUsecase {
	return &categoryUsecase{cr, cv}
}

func (cu *categoryUsecase) FindAll() ([]model.Category, error) {
	categories := []model.Category{}
	if err := cu.cr.FindAll(&categories); err != nil {
		return []model.Category{}, err
	}
	return categories, nil
}

func (cu *categoryUsecase) FindByID(id uint) (model.Category, error) {
	category := model.Category{}
	if err := cu.cr.FindByID(&category, id); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (cu *categoryUsecase) Create(category model.Category) (model.Category, error) {
	if err := cu.cv.CategoryValidate(category); err != nil {
		return model.Category{}, err
	}
	if err := cu.cr.Create(&category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (cu *categoryUsecase) Update(category model.Category) (model.Category, error) {
	if err := cu.cv.CategoryValidate(category); err != nil {
		return model.Category{}, err
	}
	if err := cu.cr.Update(&category); err != nil {
		return model.Category{}, err
	}
	return category, nil
}

func (cu *categoryUsecase) Delete(id uint) error {
	if err := cu.cr.Delete(id); err != nil {
		return err
	}
	return nil
}
