package main

import (
	"fmt"
	"geektrust/subpackage1"
	"os"
)

func main() {
	args := os.Args[1:]
	inputFile := ""
	if len(args) > 0 {
		inputFile = args[0]
	}
	fmt.Println(inputFile)
	fmt.Println(subpackage1.Tom())
}
