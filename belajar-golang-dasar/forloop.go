//for loop -> perulangan

package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
		fmt.Print(counter)
	}

	//for dg statement
	// init statement, kondisi, post statment
	for country := 1; country <= 5; country++ {
		fmt.Println("Perulang ke", country)
	}

	//for range -> untuk melakukan iterasi di array/slice/map
	slice := []string{"vini", "jumatul", "fitri"}
	//cara manual
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//cara cepat menggunakan for range
	for _, name := range slice {
		fmt.Print(" ", name)
		fmt.Println("===============")
	}

	person := make(map[string]string)
	person["name"] = "vini"
	person["title"] = "programmer"

	for key, value := range person {
		fmt.Println("key", key, "= ", value)
	}

}
