//switch expression -> sederhana dibandingkan if
//digunakan untuk melakukan pengecekan ke kondisi dalam satu varuable
package main

import "fmt"

func main() {
	name := "vini"

	switch name {
	case "eko":
		fmt.Println("hello eko")
	case "vini":
		fmt.Println("hello vini")
	default:
		fmt.Println("hi, boleh kenala")
	}

	//switch short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	//switch tanpa kondisi
	length1 := len(name)
	switch {
	case length1 > 10:
		fmt.Println("nama terlalu panjang")
	case length1 > 5:
		fmt.Println("nama terlalu panjang")
	default:
		fmt.Println("nama sudah benar")

	}
}
