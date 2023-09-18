package belajargolangjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStreamEncode(t *testing.T) {
	writer, _ := os.Create("CustomerOut.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Vini",
		LastName:   "Fitri",
		MiddleName: "Jumatul",
	}

	encoder.Encode(customer)
}
