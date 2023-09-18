//func return value -> bisa mengembalikan data
//harus menuliskan tipe data kembaliannya
// key word return
package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "helo bro"
	} else {
		return "hello" + name

	}
}

func main() {
	aku := ""
	//krn func gethalle bisa mengembalikan data kita bisa panggil fungsi nya dalam sebuah variable
	result := getHello(aku)
	fmt.Println(result)

}
