package main

import (
	"fmt"

	"github.com/regmicmahesh/crafting-interpreters/scanner"
)

func run(source []byte) {

	sc := scanner.NewScanner(string(source))

	tokens := sc.ScanTokens()

	for _, token := range tokens {
		fmt.Printf("%#v\n", token.String())
	}

}
