package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()

	router.GET("/product/:id/:status", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		id := params.ByName("id")
		text := "Product " + id + " " + params.ByName("status")
		fmt.Fprint(w, text)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/product/1/2", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Product 1 2", string(body))
}

func TestRouterCatchAllParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		image := params.ByName("image")
		fmt.Fprint(w, "Image: "+image)
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/images/small/avatar/fajar.png", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Image: /small/avatar/fajar.png", string(body))
}
