package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Vini Jumatul ", "Vini"))
	fmt.Println(strings.Contains("Vini Jumatul ", "vini"))
	fmt.Println(strings.Split("Vini/Jumatul", "/"))
	fmt.Println(strings.ToLower("Vini Jumatul"))
	fmt.Println(strings.ToUpper("Vini Jumatul"))
	fmt.Println(strings.ToTitle("Vini Jumatul"))

	fmt.Println(strings.Trim("                    Vini Jumatul   ", " ")) //ngilangin space kiri dan kanann saja

	fmt.Println(strings.ReplaceAll("Vini vini vini vini", "vini", "jum"))

}
