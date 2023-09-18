package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("Argument:")
	fmt.Println(args)

	// fmt.Println(args[0]) //merupakan lokasi compile nya
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println("hostname:", hostname)
	} else {
		fmt.Println("Error", err)
	}
}
