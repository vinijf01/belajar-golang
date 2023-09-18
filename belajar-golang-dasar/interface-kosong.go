package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	//penggunaan interface kosong contohnya dalam Println
	// var data int = Ups(1) // tidak bisa menggunakan type data int
	var data interface{} = Ups(3)
	fmt.Println(data)
}
