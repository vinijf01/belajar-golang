package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address1 := Address{"subang", "jawa barat", "indonesia"}
	address2 := &address1 //pointerss
	address3 := &address1

	address2.City = "bandung"

	*address2 = Address{"Malang", "jawa timur", "indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	address5 := &Address{"Tests", "test", "sddda"}
	fmt.Println(address5)

	//func new
	address4 := new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	alamat := Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	ChangeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
