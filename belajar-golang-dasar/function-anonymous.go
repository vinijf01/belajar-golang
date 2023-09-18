package main

import (
	"fmt"
)

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("Your are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklist1(name string) bool {
// 	return name == "admin"
// }

// func blacklist2(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("vini", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
}
