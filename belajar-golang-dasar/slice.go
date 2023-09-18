//slice -> potongan dari data array || slice reference ke array
//yang membedakannya ukuran slice bisa berubah(array fix)
//tipe data slice ->
//pointer -> penunjuk data pertama di array pada slice (pada saat buat array dulu baru slice)
//, length ->  panjang dari slice
//capacity -> kapasitas dari slice, dimana length tidak boleh lebih daru capacity
package main

import (
	"fmt"
)

func main() {

	var months = [12]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slcie1 = months[4:7]
	fmt.Println(len(slcie1))
	fmt.Println(cap(slcie1))
	fmt.Println(slcie1)

	//func slice
	//len(slice) -> untuk  mendapatkan panjang
	//cap(slice) -> untuk mendapatkan kapasitas
	//append(slice, data) -> menambah data tapi dia akan menjadi sebuah slice baru jika cap sudah penuh maka akan dibuat array baru
	//make([]TypeData,length, capacity)
	//copy(destination, source)

	var slice2 = months[10:]
	fmt.Println(slice2)

	var slice3 = append(slcie1, "vini")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"

	fmt.Println(slice3)
	fmt.Println(slice2)
	fmt.Println(months)

	//Make SLice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "vini"
	newSlice[1] = "vini1"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice : pastikan cap nya sama agar data tidak terpoting
	copySlcie := make([]string, len(newSlice), cap(newSlice))
	copy(copySlcie, newSlice)
	fmt.Println(copySlcie)

	iniArray := [...]int{1, 2, 3, 4, 5, 6}
	iniSlice := []int{1, 2, 3, 4, 5, 6}
	fmt.Print(iniArray)
	fmt.Print(iniSlice)
}
