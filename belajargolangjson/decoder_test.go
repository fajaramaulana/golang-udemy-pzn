package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

type IdentityDecoder struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func TestStreamDecoder(t *testing.T) {
	reader, err := os.Open("Identity.json")
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(reader)

	identity := &IdentityDecoder{}

	decoder.Decode(identity)

	fmt.Println(identity)

}
