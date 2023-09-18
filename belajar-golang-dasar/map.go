// tipe data map -> key-value
// key bersifat unik == index
//key word map
package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "vini",
		"address": "riau",
	}
	//namabh data
	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//func map
	//len -> mendapatkan jumlah data dalam amp
	//map[key] -> mengambil data dimap dengan key
	//map[key] = value -> mengubah data dimapa dengan key
	//make(map[typedata key]typeValue) -> membuat map baru
	//delete(map, key) -> menghapus data di map dengan key
	person["name"] = "nix"

	fmt.Println(person["name"])
	fmt.Println(person)
	delete(person, "address")
	fmt.Println(person)

}
