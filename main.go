package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

var hadError bool

func main() {

	args := os.Args[1:]

	switch len(args) {
	case 0:
		runPrompt()
	case 1:
		runFile(args[0])
	default:
		fmt.Println("Usage: golox [script]")
		os.Exit(64)
	}
}

func runFile(filePath string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	run(data)

	if hadError {
		os.Exit(65)
	}
}

func runPrompt() {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print(">")
		line := scanner.Scan()

		if !line {
			break
		}

		text := scanner.Bytes()
		run(text)
		hadError = false

	}

}
