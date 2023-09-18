//representasi dari tipe data yang sudah
//membuat tipe data alias
//key word nya type
package main

import "fmt"

func main() {
	type NoKTP string //representasi string dengan alias NoKTP
	type married bool

	var noktpvini NoKTP = "1234567890"
	fmt.Println(noktpvini)

	var marriedStatus married = true
	fmt.Println(marriedStatus)
}
