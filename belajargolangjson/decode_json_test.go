package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Identity struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func TestDecodeJson(t *testing.T) {
	jsonString := `{"first_name":"Fajar","last_name":"Agus","age":23}`
	jsonBytes := []byte(jsonString)

	identity := &Identity{}
	err := json.Unmarshal(jsonBytes, identity)
	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
	fmt.Printf("%# v\n", identity)
}
