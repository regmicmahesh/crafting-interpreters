package main

import "fmt"

func error(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where string, msg string) {

	fmt.Printf("[line %v] Error %v: %v\n", line, where, msg)
	hadError = true

}
