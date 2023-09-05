package controller

import (
	"belajarrestfullapi/helper"
	"belajarrestfullapi/model/web/request"
	"belajarrestfullapi/service"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse, err := controller.CategoryService.FindAll(request.Context())
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writer, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	helper.ReturnDataJson(writer, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse)
	return
}

func (controller *CategoryControllerImpl) FindById(writter http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)

	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	categoryResponse, err := controller.CategoryService.FindById(r.Context(), id)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	helper.ReturnDataJson(writter, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse)
	return
}

func (controller *CategoryControllerImpl) Create(writter http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryCreateRequest := request.CategoryCreateRequest{}
	err := helper.ReadFromRequestBody(r, &categoryCreateRequest)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	categoryResponse, err := controller.CategoryService.Create(r.Context(), categoryCreateRequest)

	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	helper.ReturnDataJson(writter, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse)
	return
}

func (controller *CategoryControllerImpl) Update(writter http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryUpdateRequest := request.CategoryUpdateRequest{}
	err := helper.ReadFromRequestBody(r, &categoryUpdateRequest)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	categoryUpdateRequest.Id = int(id)
	categoryResponse, err := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	helper.ReturnDataJson(writter, http.StatusOK, http.StatusText(http.StatusOK), categoryResponse)
	return
}

func (controller *CategoryControllerImpl) Delete(writter http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")

	id, err := strconv.Atoi(categoryId)
	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusBadRequest, http.StatusText(http.StatusBadRequest), err.Error())
		return
	}

	err = controller.CategoryService.Delete(r.Context(), id)

	if err != nil {
		log.Print(err.Error())
		helper.ReturnDataJson(writter, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), err.Error())
		return
	}

	helper.ReturnDataJson(writter, http.StatusOK, http.StatusText(http.StatusOK), nil)
	return
}
