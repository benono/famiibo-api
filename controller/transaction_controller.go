package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ITransactionController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type transactionController struct {
	pu usecase.ITransactionUsecase
}

func NewTransactionController(pu usecase.ITransactionUsecase) ITransactionController {
	return &transactionController{pu}
}

func (pc *transactionController) FindAll(c echo.Context) error {
	fmt.Println("Executing FindAll method in transactionController")
	transactions, err := pc.pu.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactions)
}

func (pc *transactionController) FindById(c echo.Context) error {
	id := c.Param("id")
	transactionId, _ := strconv.Atoi(id)
	transaction, err := pc.pu.FindById(uint(transactionId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transaction)
}

func (pc *transactionController) Create(c echo.Context) error {
	transaction := model.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	transactionRes, err := pc.pu.Create(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, transactionRes)
}

func (pc *transactionController) Update(c echo.Context) error {
	id := c.Param("transactionId")
	transactionId, _ := strconv.Atoi(id)
	transaction := model.Transaction{}
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	transaction.ID = uint(transactionId)
	transactionRes, err := pc.pu.Update(transaction)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, transactionRes)
}

func (pc *transactionController) Delete(c echo.Context) error {
	id := c.Param("transactionId")
	transactionId, _ := strconv.Atoi(id)
	err := pc.pu.Delete(uint(transactionId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
