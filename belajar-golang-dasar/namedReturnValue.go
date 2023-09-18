//named return  value -> membuat variable di retrun value
//k=harus menuliskna semua tipe data kemabliannya
package main

import "fmt"

func getFullname() (firstname, middlename, lastname string) {
	firstname = "vini"
	middlename = "jumatul"
	lastname = "fitri"
	return
}

func main() {
	firstname, middlename, lastname := getFullname()
	fmt.Println(firstname, middlename, lastname)
}
