package main

import (
	"fmt"
	"packagedatabase"
)

func main() {
	result := packagedatabase.GetDatabase()

	fmt.Println(result)
}
