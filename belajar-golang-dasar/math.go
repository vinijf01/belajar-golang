//operasi matematika
package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	a := 10
	b := 12
	c := a * b
	fmt.Println(c)

	//augmented assignment
	//+= -= *= /= %=
	a += 10
	fmt.Println(a)

	//unary operator
	//++(ditambah 1) --(dikurang 1) -(negative) +(positive) !(boolean kebalikan)

	a++ // a + 1
	fmt.Println(a)

	var negative = -100
	var positive = +200
	fmt.Println(negative)
	fmt.Println(positive)
}
