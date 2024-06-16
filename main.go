package main

import (
	"fmt"
	"go-rest-api/controller"
	"go-rest-api/db"
	"go-rest-api/repository"
	"go-rest-api/router"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"strconv"
)

func main() {
	db := db.NewDB()
	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()
	payeeValidator := validator.NewPayeeValidator()
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)
	payeeRepository := repository.NewPayeeRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	payeeUsecase := usecase.NewPayeeUsecase(payeeRepository, payeeValidator)
	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)
	payeeController := controller.NewPayeeController(payeeUsecase)
	e := router.NewRouter(userController, taskController, payeeController)
	e.Logger.Fatal(e.Start(":8080"))
	floatVal, _ := strconv.ParseFloat("", 64)
	fmt.Println(floatVal)
}
