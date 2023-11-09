package test

import (
	"belajarrestfullapi/app"
	"belajarrestfullapi/config"
	"belajarrestfullapi/controller"
	"belajarrestfullapi/middleware"
	"belajarrestfullapi/model/domain"
	"belajarrestfullapi/repository"
	"belajarrestfullapi/service"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func SetupRouter(db *sql.DB) http.Handler {

	validation := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validation)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func ExecTruncate(db *sql.DB) {
	_, err := db.Exec("TRUNCATE TABLE categories")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestCreateCategorySuccess(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name":"Charger Laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/categories", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, 201, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 201, int(responseBody["code"].(float64)))
	assert.Equal(t, "Created", responseBody["status"])
	assert.Equal(t, "Charger Laptop", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, time.Now().Format(time.RFC3339), responseBody["data"].(map[string]interface{})["created_at"])
	assert.Equal(t, time.Now().Format(time.RFC3339), responseBody["data"].(map[string]interface{})["updated_at"])
}

func TestCreateCategoryWithoutBodyRequest(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/categories", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), responseBody["status"])
}

func TestCreateCategoryFailedValidation(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8081/api/categories", requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "error validation", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	requestBody := strings.NewReader(`{"name":"Charger Laptop Updated"}`)
	id := strconv.Itoa(category.Id)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8081/api/categories/"+id, requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
	assert.Equal(t, "Charger Laptop Updated", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, time.Now().Format(time.RFC3339), responseBody["data"].(map[string]interface{})["updated_at"])
	assert.Equal(t, category.CreatedAt.Format(time.RFC3339), responseBody["data"].(map[string]interface{})["created_at"])
}

func TestUpdatedCategoryWithoutBodyRequest(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	id := strconv.Itoa(category.Id)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8081/api/categories/"+id, nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), responseBody["status"])
}

func TestUpdateCategoryFailedValidation(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	id := strconv.Itoa(category.Id)
	requestBody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8081/api/categories/"+id, requestBody)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "error validation", responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	id := strconv.Itoa(category.Id)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories/"+id, nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
	assert.Equal(t, "Charger Laptop", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["id"])
}

func TestGetCategoryFailed(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories/1", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	id := strconv.Itoa(category.Id)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8081/api/categories/"+id, nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	fmt.Printf("%# v\n", responseBody)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	ExecTruncate(db)
	router := SetupRouter(db)

	// create category first
	tx, err := db.Begin()

	if err != nil {
		log.Fatal(err.Error())
	}
	categoryRepository := repository.NewCategoryRepository()
	_, err = categoryRepository.Save(context.Background(), tx, domain.Category{
		Name:      "Charger Laptop",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8081/api/categories/2", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	fmt.Printf("%# v\n", responseBody)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestGetAllCategorySuccess(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	fmt.Printf("%# v\n", responseBody)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusOK), responseBody["status"])
	assert.Equal(t, 1, int(responseBody["data"].([]interface{})[0].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "Charger Laptop", responseBody["data"].([]interface{})[0].(map[string]interface{})["name"])
}

func TestGetAllCategoryFailed(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasialah")

	recorder := httptest.NewRecorder()

	db.Close()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	fmt.Printf("%# v\n", responseBody)

	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)
	assert.Equal(t, http.StatusInternalServerError, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusInternalServerError), responseBody["status"])
}

func TestUnauthorized(t *testing.T) {
	db, err := config.NewDBTest()
	if err != nil {
		log.Fatal(err.Error())
	}
	router := SetupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8081/api/categories", nil)

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Api-Key", "Rahasia")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	var responseBody map[string]interface{}

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err.Error())
	}

	json.Unmarshal(body, &responseBody)

	fmt.Printf("%# v\n", responseBody)

	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)
	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), responseBody["status"])
}
