package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	nashir := Man{"Nashir"}
	nashir.Married()

	fmt.Println(nashir.Name)
}
