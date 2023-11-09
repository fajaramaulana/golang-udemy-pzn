package main

import (
	"belajarrestfullapi/app"
	"belajarrestfullapi/config"
	"belajarrestfullapi/controller"
	"belajarrestfullapi/middleware"
	"belajarrestfullapi/repository"
	"belajarrestfullapi/service"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db, err := config.NewDB()
	if err != nil {
		log.Print(err.Error())
	}

	validation := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validation)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "127.0.0.1:8081",
		Handler: middleware.NewAuthMiddleware(router),
	}

	log.Println("Server running on port 8081")
	fmt.Println("Server running on port 8081")

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err.Error())
	}
}
