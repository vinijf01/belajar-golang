package main

import "fmt"

func main() {
	counter := 0
	name := "Eko"

	increment := func() {
		name := "Vini"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
