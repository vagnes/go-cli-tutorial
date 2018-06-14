package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {

	flag.Usage = func() {
		fmt.Println("\nMy Mediocre Program - By Me\n")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	var operator string
	flag.StringVar(&operator, "o", "+", "operator to use (+, -, * or /)")

	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for i := 0; i < len(flag.Args()); i = i + 2 {
		firstInt, _ := strconv.ParseFloat(flag.Arg(i), 64)
		secondInt, _ := strconv.ParseFloat(flag.Arg(i+1), 64)
		if operator == "+" {
			fmt.Printf("%.2f + %.2f = %.2f\n", firstInt, secondInt, firstInt+secondInt)
		} else if operator == "-" {
			fmt.Printf("%.2f - %.2f = %.2f\n", firstInt, secondInt, firstInt-secondInt)
		} else if operator == "*" {
			fmt.Printf("%.2f * %.2f = %.2f\n", firstInt, secondInt, firstInt*secondInt)
		} else if operator == "/" {
			fmt.Printf("%.2f / %.2f = %.2f\n", firstInt, secondInt, firstInt/secondInt)
		} else {
			println("Operator not recognised.")
		}
	}

}
