package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Halo", hasName.GetName())
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
	var nashir Person
	nashir.Name = "Nashir"
	SayHello(nashir)

	var cat Animal
	cat.Name = "Meong"
	SayHello(cat)

}
