package main

import "fmt"

//dengan recover: panic tidak akan menghentikan programnya
// recover disimpan di defer function
func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Applikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Vini")
}
