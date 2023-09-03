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

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	router.POST("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "INI POST")
	})

	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ups, something went wrong. %s", http.StatusText(http.StatusMethodNotAllowed))
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Ups, something went wrong. "+http.StatusText(http.StatusMethodNotAllowed), string(bytes))
}
