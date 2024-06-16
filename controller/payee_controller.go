package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IPayeeController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type payeeController struct {
	pu usecase.IPayeeUsecase
}

func NewPayeeController(pu usecase.IPayeeUsecase) IPayeeController {
	return &payeeController{pu}
}

func (pc *payeeController) FindAll(c echo.Context) error {
	fmt.Println("Executing FindAll method in payeeController")
	payees, err := pc.pu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payees)
}

func (pc *payeeController) FindById(c echo.Context) error {
	id := c.Param("id")
	payeeId, _ := strconv.Atoi(id)
	payee, err := pc.pu.FindById(uint(payeeId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payee)
}

func (pc *payeeController) Create(c echo.Context) error {
	payee := model.Payee{}
	if err := c.Bind(&payee); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	payeeRes, err := pc.pu.Create(payee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, payeeRes)
}

func (pc *payeeController) Update(c echo.Context) error {
	id := c.Param("payeeId")
	payeeId, _ := strconv.Atoi(id)
	payee := model.Payee{}
	if err := c.Bind(&payee); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	payee.ID = uint(payeeId)
	payeeRes, err := pc.pu.Update(payee)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, payeeRes)
}

func (pc *payeeController) Delete(c echo.Context) error {
	id := c.Param("payeeId")
	payeeId, _ := strconv.Atoi(id)
	err := pc.pu.Delete(uint(payeeId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
