package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	//CARA 1
	// var vini Customer
	// vini.Name = "Vini"
	// vini.Address = "Indonesia"
	// vini.Age = 22

	// fmt.Println(vini)
	// fmt.Println(vini.Name)

	//CARA 2
	// vini := Customer{
	// 	Name:    "vini",
	// 	Address: "Indonesia",
	// 	Age:     22,
	// }

	//CARA 3 (tidak direkomendasikan krn rentan error)
	vini := Customer{"vini", "indonesia", 22}
	fmt.Println(vini)

}
