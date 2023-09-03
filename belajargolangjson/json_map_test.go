package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"first_name":"Fajar","last_name":"Agus","age":23}`
	jsonBytes := []byte(jsonString)

	var identity map[string]interface{}

	err := json.Unmarshal(jsonBytes, &identity)

	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
	fmt.Println(identity["first_name"])
}

func TestMapEncode(t *testing.T) {
	identity := map[string]interface{}{
		"first_name": "Fajar",
		"last_name":  "Agus",
		"age":        23,
	}

	bytes, err := json.Marshal(identity)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}
