package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type IdentityJsonObject struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func TestJSONObject(t *testing.T) {
	identity := IdentityJsonObject{
		FirstName: "Fajar",
		LastName:  "Agus",
		Age:       23,
	}

	bytes, err := json.Marshal(identity)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
