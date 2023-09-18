package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err)
	}
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err)
	}

	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	valueINt, err := strconv.Atoi("2000000")
	if err == nil {
		fmt.Println(valueINt)
	}
}
