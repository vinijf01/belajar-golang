//tipe data array -> data yang berisikikan kumpulan data dengan tipe yang sama
//saat membuat array jumlah data yang mesti ditampung array harus ditentukan
//daya tampung array tidak bisa ditambahkan setelah array dibuat
//index dimulai dari 0, kita bisa akses data menggunakan index
package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "vini"
	name[1] = "jumatul"
	name[2] = "fitri"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	//mem array langsung saat deklarasi
	var value = [3]int{5, 4, 3}
	fmt.Println(value[2])

	//func array
	//len -> menampilkan panjang array
	//array[index] -> mendapatkan data di posisi index
	//array[index] = value -> mengubah data di posisi index

	value[2] = 89
	fmt.Println(value)
	fmt.Println(len(value))

	var angka = []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(angka)

}
