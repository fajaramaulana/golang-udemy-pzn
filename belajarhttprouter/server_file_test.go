package main

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	router := httprouter.New()
	directory, err := fs.Sub(resources, "resources")
	if err != nil {
		panic(err)
	}
	router.ServeFiles("/static/*filepath", http.FS(directory))

	request := httptest.NewRequest("GET", "http://localhost:8080/static/hi.txt", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Hi from fajar", string(body))
}
