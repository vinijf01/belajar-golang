package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	var contohError error = errors.New("upss Error")
	fmt.Println(contohError)

	// Pembagi(4, 2)
	// fmt.Println(Pembagi(4, 2))

	hasil, err := Pembagi(100, 0)
	if err == nil {
		fmt.Println("Hasil: ", hasil)
	} else {
		fmt.Println("Error: ", err.Error())
	}
}
