package service

import (
	"belajarrestfullapi/helper"
	"belajarrestfullapi/model/domain"
	"belajarrestfullapi/model/web/request"
	"belajarrestfullapi/model/web/response"
	"belajarrestfullapi/repository"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type CategoryService interface {
	FindAll(ctx context.Context) ([]response.CategoryResponseAll, error)
	FindById(ctx context.Context, categoryId int) (response.CategoryResponseById, error)
	Create(ctx context.Context, request request.CategoryCreateRequest) (response.CategoryResponse, error)
	Update(ctx context.Context, request request.CategoryUpdateRequest) (response.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int) error
}

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validation         *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validation *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validation:         validation,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) ([]response.CategoryResponseAll, error) {
	tx, err := service.DB.Begin()
	if err != nil {
		return nil, err
	}

	// defer digunakan untuk mengeksekusi kode di dalamnya ketika fungsi sudah selesai dijalankan
	// defer akan selalu dijalankan walaupun terjadi panic
	defer helper.CommitOrRollback(tx)

	categories, err := service.CategoryRepository.FindAll(ctx, tx)

	if err != nil {
		return nil, err
	}
	categoryResponses := []response.CategoryResponseAll{}
	for _, category := range categories {
		categoryResponses = append(categoryResponses, response.CategoryResponseAll{
			Id:   category.Id,
			Name: category.Name,
		})
	}

	return categoryResponses, nil
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) (response.CategoryResponseById, error) {
	tx, err := service.DB.Begin()
	categoryResponseEmpty := response.CategoryResponseById{}
	if err != nil {
		return categoryResponseEmpty, err
	}

	// defer digunakan untuk mengeksekusi kode di dalamnya ketika fungsi sudah selesai dijalankan
	// defer akan selalu dijalankan walaupun terjadi panic
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)

	if err != nil {
		return categoryResponseEmpty, err
	}

	return response.CategoryResponseById{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request request.CategoryCreateRequest) (response.CategoryResponse, error) {
	categoryResponseEmpty := response.CategoryResponse{}
	err := service.Validation.Struct(request)

	if err != nil {
		return categoryResponseEmpty, err
	}
	tx, err := service.DB.Begin()

	if err != nil {
		return categoryResponseEmpty, err
	}

	// defer digunakan untuk mengeksekusi kode di dalamnya ketika fungsi sudah selesai dijalankan
	// defer akan selalu dijalankan walaupun terjadi panic
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name:      request.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	category, err = service.CategoryRepository.Save(ctx, tx, category)

	if err != nil {
		return categoryResponseEmpty, err
	}

	return response.CategoryResponse{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt.Format(time.RFC3339),
		UpdatedAt: category.UpdatedAt.Format(time.RFC3339),
	}, nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request request.CategoryUpdateRequest) (response.CategoryResponse, error) {
	categoryResponseEmpty := response.CategoryResponse{}
	err := service.Validation.Struct(request)
	if err != nil {
		return categoryResponseEmpty, err
	}
	tx, err := service.DB.Begin()
	if err != nil {
		return categoryResponseEmpty, err
	}

	// defer digunakan untuk mengeksekusi kode di dalamnya ketika fungsi sudah selesai dijalankan
	// defer akan selalu dijalankan walaupun terjadi panic
	defer helper.CommitOrRollback(tx)

	dataOld, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		return categoryResponseEmpty, err
	}

	categoryUpdate := domain.Category{
		Id:        request.Id,
		Name:      request.Name,
		CreatedAt: dataOld.CreatedAt,
		UpdatedAt: time.Now(),
	}

	categoryUpdate, err = service.CategoryRepository.Update(ctx, tx, categoryUpdate)
	if err != nil {
		return categoryResponseEmpty, err
	}

	return response.CategoryResponse{
		Id:        categoryUpdate.Id,
		Name:      categoryUpdate.Name,
		CreatedAt: categoryUpdate.CreatedAt.Format(time.RFC3339),
		UpdatedAt: categoryUpdate.UpdatedAt.Format(time.RFC3339),
	}, nil

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) error {
	tx, err := service.DB.Begin()
	if err != nil {
		return err
	}

	// defer digunakan untuk mengeksekusi kode di dalamnya ketika fungsi sudah selesai dijalankan
	// defer akan selalu dijalankan walaupun terjadi panic
	defer helper.CommitOrRollback(tx)

	_, err = service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		return err
	}

	err = service.CategoryRepository.Delete(ctx, tx, categoryId)

	if err != nil {
		return err
	}

	return nil
}
