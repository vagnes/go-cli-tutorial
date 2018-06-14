package main

import (
	"fmt"
	"os"
)

func main() {
	
	for i, arg := range os.Args[1:] { // 1, not 0 because we do not want the name of the program
        fmt.Printf("Argument No.%d: %s\n", i+1, arg)
    }
}