//variable yang nilainya tidak bisa diubah lagi, key word "const"
//wajib menginisialisasikan nilai nya langsung
package main

import "fmt"

func main() {
	const firstname string = "eko"
	const lastname = "fitri"
	const value = 1000

	fmt.Println(firstname)

	//multiple constant
	const (
		firstName string = "vini"
		lastName         = "jumatul"
		valuea           = 100
	)
}
