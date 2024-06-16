package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ICategoryController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type categoryController struct {
	cu usecase.ICategoryUsecase
}

func NewCategoryController(cu usecase.ICategoryUsecase) ICategoryController {
	return &categoryController{cu}
}

func (cc *categoryController) FindAll(c echo.Context) error {
	categories, err := cc.cu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categories)
}

func (cc *categoryController) FindById(c echo.Context) error {
	id := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(id)
	category, err := cc.cu.FindByID(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, category)
}

func (cc *categoryController) Create(c echo.Context) error {
	category := model.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	categoryRes, err := cc.cu.Create(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, categoryRes)
}

func (cc *categoryController) Update(c echo.Context) error {
	id := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(id)
	category := model.Category{}
	if err := c.Bind(&category); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	category.ID = uint(categoryId)
	categoryRes, err := cc.cu.Update(category)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, categoryRes)
}

func (cc *categoryController) Delete(c echo.Context) error {
	id := c.Param("categoryId")
	categoryId, _ := strconv.Atoi(id)
	err := cc.cu.Delete(uint(categoryId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
