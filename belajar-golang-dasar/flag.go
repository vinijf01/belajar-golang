package main

import (
	"flag"
	"fmt"
)

// go run flag.go -host=127.0.0.0 -user=vini -password=admin123
func main() {
	var host *string = flag.String("host", "localhost", "Put your host") // nama nilaiDefault deskripsi
	var user *string = flag.String("user", "root", "Put your database use")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 100, "put your number")
	flag.Parse() // after all flags are defined, call this

	fmt.Println("host :", *host)
	fmt.Println("user :", *user)
	fmt.Println("password :", *password)
	fmt.Println("number :", *number)
}
