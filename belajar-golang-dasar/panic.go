package main

import "fmt"

func endApp() {
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
}
