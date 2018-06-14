package main

import (
	"fmt"
	"os"
)

func main() {
    for i := 0; i<5; i++ {
		programName := os.Args[i]
		fmt.Printf("Argument No. : %d (%s)\n", i, programName)
	}
    
}