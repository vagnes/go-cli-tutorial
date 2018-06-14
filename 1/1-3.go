package main

import (
	"fmt"
	"os"
)

func main() {
	numberOfArgs := len(os.Args[1:])
    for i := 0; i <= numberOfArgs; i++ {
		programName := os.Args[i]
		fmt.Printf("Argument No. : %d (%s)\n", i, programName)
	}
    
}