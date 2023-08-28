package belajargolangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	allName := r.URL.Query()["name"]
	// get second value inside variable allName
	fmt.Printf("%# v\n", allName[1])
	// loop all name
	for _, each := range allName {
		fmt.Println(each)
	}
	// fmt.Printf("%# v\n", r.URL.Query().Get("name"))
	if name == "" {
		name = "World"
	} else {
		name = name + "!"
	}

	fmt.Fprintf(w, "Hello %s", name)
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/?name=Fajar&name=tangerang", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	recorderResult := recorder.Result()
	body, _ := io.ReadAll(recorderResult.Body)
	fmt.Println(string(body))
}
