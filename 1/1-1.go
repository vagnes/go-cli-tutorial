package main

import (
	"fmt"
	"os"
)

func main() {
    programName := os.Args[0]
    fmt.Printf("Program name: %s\n", programName)
}