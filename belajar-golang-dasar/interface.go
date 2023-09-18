package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHel(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var vini Person
	vini.Name = "RARA"

	SayHel(vini)

	cat := Animal{
		Name: "Meow",
	}

	SayHel(cat)
}
