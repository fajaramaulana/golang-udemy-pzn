package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type IdentityJsonArray struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       int       `json:"age"`
	Hobbies   []string  `json:"hobbies"`
	Addresses []Address `json:"addresses"`
}

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

func TestEncodeJsonArray(t *testing.T) {
	identity := IdentityJsonArray{
		FirstName: "Fajar",
		LastName:  "Agus",
		Age:       23,
		Hobbies:   []string{"Reading", "Coding", "Gaming"},
	}

	bytes, err := json.Marshal(identity)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}

func TestDecodeJsonArray(t *testing.T) {
	jsonString := `{"first_name":"Fajar","last_name":"Agus","age":23,"hobbies":["Reading","Coding","Gaming"]}`
	jsonBytes := []byte(jsonString)

	identity := &IdentityJsonArray{}
	err := json.Unmarshal(jsonBytes, identity)
	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
	fmt.Printf("%# v\n", identity)
}

func TestJsonArrayComplexEncode(t *testing.T) {
	identity := IdentityJsonArray{
		FirstName: "Fajar",
		LastName:  "Agus",
		Age:       23,
		Hobbies:   []string{"Reading", "Coding", "Gaming"},
		Addresses: []Address{
			{
				Street:     "Jl. Raya",
				City:       "Jakarta",
				PostalCode: "12345",
			},
			{
				Street:     "Jl. Raya2",
				City:       "Jakarta2",
				PostalCode: "123456",
			},
		},
	}

	bytes, err := json.Marshal(identity)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"first_name":"Fajar","last_name":"Agus","age":23,"hobbies":["Reading","Coding","Gaming"],"addresses":[{"street":"Jl. Raya","city":"Jakarta","postal_code":"12345"},{"street":"Jl. Raya2","city":"Jakarta2","postal_code":"123456"}]}`
	jsonBytes := []byte(jsonString)

	identity := &IdentityJsonArray{}
	err := json.Unmarshal(jsonBytes, identity)
	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
	fmt.Printf("%# v\n", identity)
}
