package main

import (
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
)

func main() {
	fmt.Println("Starting Server....ç")
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	accountValidator := validator.NewAccountValidator()
	payeeValidator := validator.NewPayeeValidator()
	categoryValidator := validator.NewCategoryValidator()
	userRepository := repository.NewUserRepository(db)
	accountRepository := repository.NewAccountRepository(db)
	payeeRepository := repository.NewPayeeRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	accountUsecase := usecase.NewAccountUsecase(accountRepository, accountValidator)
	payeeUsecase := usecase.NewPayeeUsecase(payeeRepository, payeeValidator)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepository, categoryValidator)
	userController := controller.NewUserController(userUsecase)
	accountController := controller.NewAccountController(accountUsecase)
	payeeController := controller.NewPayeeController(payeeUsecase)
	categoryController := controller.NewCategoryController(categoryUsecase)
	e := router.NewRouter(userController, payeeController, accountController, categoryController)
	e.Logger.Fatal(e.Start(":8080"))
}
