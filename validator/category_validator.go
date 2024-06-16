package validator

import (
	"go-rest-api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ICategoryValidator interface {
	CategoryValidate(category model.Category) error
}

type categoryValidator struct{}

func NewCategoryValidator() ICategoryValidator {
	return &categoryValidator{}
}

func (cv *categoryValidator) CategoryValidate(category model.Category) error {
	return validation.ValidateStruct(&category,
		validation.Field(
			&category.Name,
			validation.Required.Error("name is required"),
		),
	)
}
