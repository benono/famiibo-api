package controller

import (
	"fmt"
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type IAccountController interface {
	FindAll(c echo.Context) error
	FindById(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type accountController struct {
	au usecase.IAccountUsecase
}

func NewAccountController(au usecase.IAccountUsecase) IAccountController {
	return &accountController{au}
}

func (ac *accountController) FindAll(c echo.Context) error {
	fmt.Println("Executing FindAll method in accountController")
	accounts, err := ac.au.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, accounts)
}

func (pc *accountController) FindById(c echo.Context) error {
	id := c.Param("id")
	accountId, _ := strconv.Atoi(id)
	account, err := pc.au.FindById(uint(accountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, account)
}

func (pc *accountController) Create(c echo.Context) error {
	account := model.Account{}
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	accountRes, err := pc.au.Create(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, accountRes)
}

func (pc *accountController) Update(c echo.Context) error {
	id := c.Param("accountId")
	accountId, _ := strconv.Atoi(id)
	account := model.Account{}
	if err := c.Bind(&account); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	account.ID = uint(accountId)
	accountRes, err := pc.au.Update(account)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, accountRes)
}

func (pc *accountController) Delete(c echo.Context) error {
	id := c.Param("accountId")
	accountId, _ := strconv.Atoi(id)
	err := pc.au.Delete(uint(accountId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
