package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Ayam" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Vini", spamFilter)
	sayHelloWithFilter("Ayam", spamFilter)
}
