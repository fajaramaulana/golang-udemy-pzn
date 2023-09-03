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

func TestPanic(t *testing.T) {
	router := httprouter.New()

	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		textPanic := "Ups, something went wrong. "
		fmt.Fprint(w, textPanic, error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		panic("ups")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Ups, something went wrong. ups", string(bytes))
}
