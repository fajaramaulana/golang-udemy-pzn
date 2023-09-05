package middleware

import (
	"belajarrestfullapi/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writter http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "Rahasialah" {
		middleware.Handler.ServeHTTP(writter, request)
	} else {
		helper.ReturnDataJson(writter, http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), "Unauthorized")
	}
}
