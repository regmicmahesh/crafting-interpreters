package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/regmicmahesh/crafting-interpreters/scanner"
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

func error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {

	fmt.Printf("[line %v] Error %v: %v\n", line, where, msg)
	hadError = true

}

func run(source []byte) {

	sc := scanner.NewScanner(string(source))

	tokens := sc.ScanTokens()

	for _, token := range tokens {
		fmt.Printf("%#v\n", token.String())
	}

}
