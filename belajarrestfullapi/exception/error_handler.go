package exception

import (
	"belajarrestfullapi/model/web/response"
	"encoding/json"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	response := response.ResponseReturn{
		Code:   http.StatusInternalServerError,
		Status: http.StatusText(http.StatusInternalServerError),
		Data:   err.(error).Error(),
	}

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}
