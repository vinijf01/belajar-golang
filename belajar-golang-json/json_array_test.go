package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	customer := Customer{
		FirstName:  "Vini",
		MiddleName: "Jumatul",
		LastName:   "Fitri",
		Age:        22,
		Hobbies:    []string{"coding", "baca", "game"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonstring := `{"FirstName":"Vini","MiddleName":"Jumatul","LastName":"Fitri","Age":22,"Hobbies":["coding","baca","game"]}`
	jsonBytes := []byte(jsonstring)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName: "Vini",
		Addresses: []Address{
			{
				Street:     "Jalan Ayam",
				Country:    "Hewan",
				PostalCode: "A1",
			},
			{
				Street:     "Jalan Kucing",
				Country:    "Hewan",
				PostalCode: "K1",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonstring := `{"FirstName":"Vini","MiddleName":"","LastName":"","Age":0,"Married":false,"Hobbies":null,"Addresses":[{"Street":"Jalan Ayam","Country":"Hewan","PostalCode":"A1"},{"Street":"Jalan Kucing","Country":"Hewan","PostalCode":"K1"}]}`
	jsonBytes := []byte(jsonstring)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonstring := `[{"Street":"Jalan Ayam","Country":"Hewan","PostalCode":"A1"},{"Street":"Jalan Kucing","Country":"Hewan","PostalCode":"K1"}]`
	jsonBytes := []byte(jsonstring)

	addresses := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresses)
	if err != nil {
		panic(err)
	}
	fmt.Println(addresses)
}

//atau

func TestOnlyJSONArrayComplex(t *testing.T) {
	addresses := []Address{
		{
			Street:     "Jalan Ayam",
			Country:    "Hewan",
			PostalCode: "A1",
		},
		{
			Street:     "Jalan Kucing",
			Country:    "Hewan",
			PostalCode: "K1",
		},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}
