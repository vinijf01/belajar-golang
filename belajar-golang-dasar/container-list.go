package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Vini")
	data.PushBack("Jumatul")
	data.PushBack("Fitri")
	data.PushFront("Nici")

	// fmt.Println(data.Front().Next().Value)
	// fmt.Println(data.Back().Value)

	// for element := data.Front(); element != nil; element = element.Next() {
	// 	fmt.Println(element.Value)
	// }

	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
