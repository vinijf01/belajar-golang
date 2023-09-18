package main

import (
	// _ "belajar-golang-dasar/database" // _ disebut blank identifier
	"belajar-golang-dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

}
