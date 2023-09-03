package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type IdentityEncoder struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func TestEncoder(t *testing.T) {
	writer, err := os.Create("output_identity.json")

	if err != nil {
		panic(err)
	}

	identity := IdentityEncoder{
		FirstName: "Fajar",
		LastName:  "Agus Maulana",
		Age:       23,
	}

	encoder := json.NewEncoder(writer)

	err = encoder.Encode(identity)

	if err != nil {
		panic(err)
	}

	fmt.Println(identity)
}
