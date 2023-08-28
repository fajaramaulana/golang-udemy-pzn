package belajargolangweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func WriteCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Powered-By"
	cookie.Value = "Fajarr"
	cookie.Path = "/"

	cookie1 := new(http.Cookie)
	cookie1.Name = "X-Powered-By"
	cookie1.Value = "Agus"
	cookie1.Path = "/"

	http.SetCookie(writer, cookie)
	http.SetCookie(writer, cookie1)
	fmt.Fprint(writer, "add cookie success")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Powered-By")

	if err != nil {
		fmt.Fprint(writer, "no cookie")
	} else {
		fmt.Fprint(writer, cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/write-cookie", WriteCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandleCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/write-cookie", nil)

	recorder := httptest.NewRecorder()

	WriteCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s: %s\n", cookie.Name, cookie.Value)
	}
}
