//return multiple value -> lebih dari satu kembalian nilai/data
//k=harus menuliskna semua tipe data kemabliannya
package main

import "fmt"

func getFullname() (string, string) {
	return "vini", "jumatul"
}

func main() {
	firstname, lastname := getFullname()
	fmt.Println(firstname, lastname)
}
