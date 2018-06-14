package main

import (
	"flag"
	"fmt"
)

func main() {
	
	var password string
	flag.StringVar(&password, "p", "", "password for access")
	flag.Parse()

	if password == "password" {
		fmt.Println("Access granted")
	} else {
		fmt.Println("Access denied")
	}

}