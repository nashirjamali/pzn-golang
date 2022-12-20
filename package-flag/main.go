package main

import (
	"flag"
	"fmt"
)

func main() {
	// go run package-flag/main.go -host=localhost -user=root
	// go run package-flag/main.go
	var hostname *string = flag.String("host", "localhost", "Put your database host name")
	var user *string = flag.String("user", "root", "Put your database user name")

	flag.Parse()

	fmt.Println("Host :", *hostname)
	fmt.Println("User :", *user)
}
