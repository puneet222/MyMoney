package main

import (
	"fmt"
	"geektrust/commander"
	"geektrust/portfolio"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	inputFile := "input.txt"
	if len(args) > 0 {
		inputFile = args[0]
	}
	// read input file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error while reading file", err)
	}
	// generate commands from input file
	commands := commander.GenerateCommands(data)
	// generate portfolio from commands
	startYear := time.Now().Year()
	_ = portfolio.BuildPortfolio(commands, startYear)
	// for printing portfolio
	//fmt.Println(p)
}
