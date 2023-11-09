package app

import (
	"belajarrestfullapi/controller"
	"belajarrestfullapi/exception"
	"belajarrestfullapi/helper"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
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

	return router
}
