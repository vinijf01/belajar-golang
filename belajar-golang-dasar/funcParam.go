//func parameter -> data dari luar func ini disebut param
//param boleh lebih dari satu, param tidak wajib
// jika ada param di func kita wajib memasukan data ke dalam param ketika memanggil func nya

package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func main() {
	sayHello("vini")

	ass := "aku"
	sayHello(ass)

}
