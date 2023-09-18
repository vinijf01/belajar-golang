package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street, Country, PostalCode string
}

type Customer struct {
	FirstName, MiddleName, LastName string
	Age                             int
	Married                         bool
	Hobbies                         []string
	Addresses                       []Address
}

func TestObject(t *testing.T) {
	customer := Customer{
		FirstName:  "Vini",
		MiddleName: "Jumatul",
		LastName:   "Fitri",
		Age:        22,
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}
