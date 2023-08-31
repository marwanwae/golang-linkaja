package main

import (
	"belajar-go-rest-api/app"
	"belajar-go-rest-api/configs"
	"belajar-go-rest-api/controller"
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/middleware"
	"belajar-go-rest-api/repository"
	"belajar-go-rest-api/services"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func init() {
	configs.LoadConfig(".")
}

func main() {
	db := app.NewDB()
	validate := validator.New()
	AccountRepository := repository.NewAccountRepository()
	AccountService := services.NewAccountService(AccountRepository, db, validate)
	AccountController := controller.NewAccountController(AccountService)

	router := app.NewRouter(AccountController)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: middleware.NewAuthMiddleware(router),
	}

	fmt.Println("server is runing localhost:8081")

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
