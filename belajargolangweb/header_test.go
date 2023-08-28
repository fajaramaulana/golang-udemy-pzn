package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func ResponseWriter(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Powered-By", "Fajar")
	fmt.Fprint(w, "OK")
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)

	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	recorder := httptest.NewRecorder()

	ResponseWriter(recorder, request)

	response := recorder.Result()

	fmt.Println(response.Header)

	fmt.Println(response.Header.Get("X-Powered-By"))
}
