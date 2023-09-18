//if expression -> untuk percabangan
//kita bisa mengeksekusi suatu kode program jika suatu konsisi terpenuhi

package main

import "fmt"

func main() {
	name := "vini1"

	if name == "vini" {
		fmt.Println("hello vini")
	} else if name == "vini1" {
		fmt.Println("hello vini1")

	} else {
		fmt.Println("hello, boleh kenalan?")
	}

	//if dg short statement
	if length := len(name); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")

	}
}
