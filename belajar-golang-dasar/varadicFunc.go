//variadic func ->
//vararg (variable argument) -> cirinya ...,  diletakan paling belakang, bisa memasukan banyak data
package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//slice sbg parameter caranya tambahakan ...

func main() {
	total := sumAll(10, 2, 1, 6, 5, 8)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	total = sumAll(slice...)
	fmt.Println(total)

}
