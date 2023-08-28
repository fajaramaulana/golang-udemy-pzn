package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, "Name is empty", http.StatusBadRequest)
		return
	} else {
		w.WriteHeader(http.StatusOK)
		http.Error(w, "OK", http.StatusOK)
		return
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?name=fajar", nil)

	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
