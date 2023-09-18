//membandingkan dua buah data
//hasilnya adalah nilai boolean
//> < >= <= == !=

package main

import "fmt"

func main() {
	name1 := "vini"
	name2 := "vini"

	result := name1 == name2
	fmt.Println(result)

	result = name1 != name2
	fmt.Println(result)

	value1 := 100
	value2 := 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
}
