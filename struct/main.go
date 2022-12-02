package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

// Struct method
func (customer Customer) sayHello() {
	fmt.Println("Hello, nama saya " + customer.Name)
}

func main() {
	var nashir Customer
	nashir.Name = "Nashir Jamali"
	nashir.Address = "Indonesia"
	nashir.Age = 23

	fmt.Println(nashir)
	fmt.Println(nashir.Name)

	// Cara kedua, struct literal
	azzu := Customer{
		Name:    "Azzukhruf",
		Address: "Indonesia",
		Age:     25,
	}

	fmt.Println(azzu)

	// Struct method
	nashir.sayHello()
}
