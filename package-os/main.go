package main

import (
	"fmt"
	"os"
)

func main() {

	// Args
	args := os.Args
	fmt.Println(args)

	// go run package-os/main.go argumen1 argumen2 argumen3
	fmt.Println(args[1])
	fmt.Println(args[2])
	fmt.Println(args[3])

	// Hostname
	name, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(name)
	}

	// Get Environment

	// export APP_USERNAME="nashirjamali"
	username := os.Getenv("APP_USERNAME")
	// export APP_PASSWORD="rahasia"
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("Username :", username)
	fmt.Println("Password :", password)

}
