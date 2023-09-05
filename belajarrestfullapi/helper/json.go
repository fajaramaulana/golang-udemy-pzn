package helper

import (
	"belajarrestfullapi/model/web/response"
	"encoding/json"
	"net/http"
)

func ReturnDataJson(writter http.ResponseWriter, code int, status string, data interface{}) {
	writter.Header().Set("Content-Type", "application/json")
	response := response.ResponseReturn{
		Code:   code,
		Status: status,
		Data:   data,
	}

	if data == "no data found" {
		response.Code = http.StatusNotFound
		response.Status = http.StatusText(http.StatusNotFound)
		response.Data = "no data found"
	}

	writter.WriteHeader(response.Code)

	encoder := json.NewEncoder(writter)
	err := encoder.Encode(response)
	if err != nil {
		panic(err)
	}
}

func ReadFromRequestBody(request *http.Request, data interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(data)
	if err != nil {
		return err
	}
	return nil
}
