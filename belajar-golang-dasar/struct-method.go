package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// func sayHI(customer Customer, name string) {
// 	fmt.Println("Hello", name, "nama sayaa ", customer.Name)
// }

func (customer Customer) sayHI(name string) {
	fmt.Println("Hallo", name, "Nama Saya", customer.Name)
}

func main() {
	eko := Customer{
		Name:    "Vini",
		Address: "Indonesia",
		Age:     22,
	}

	// sayHI(eko, "abi")
	eko.sayHI("abi")
}
