package main

import (
	"belajarrestfullapi/config"
	"belajarrestfullapi/controller"
	"belajarrestfullapi/exception"
	"belajarrestfullapi/helper"
	"belajarrestfullapi/middleware"
	"belajarrestfullapi/repository"
	"belajarrestfullapi/service"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ReturnDataJson(w, http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
		return
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		helper.ReturnDataJson(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed), nil)
		return
	})

	router.PanicHandler = exception.ErrorHandler

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
