package main

import "fmt"

// interface kosong = interface apapun bisa diisi tipe data apapun
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var data interface{} = Ups(2)
	fmt.Println(data)
}
