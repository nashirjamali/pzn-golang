package main

import "fmt"

func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string)

	fmt.Println(resultString)

	// Type assertion menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown type")
	}
}
