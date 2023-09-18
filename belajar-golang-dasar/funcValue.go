//func -> first class citizen, bisa dibaut sbg tipe data, bisa dismpan dalam sebuah variabel
//func value = funcgsi sebagai value artinya kita simpan ke dalam sebuah varaible
package main

import "fmt"

func getGoodbye(name string) string {
	return "Good by " + name
}
func main() {
	sayGoodBye := getGoodbye

	result := sayGoodBye("Vini")
	fmt.Println(result)
}
