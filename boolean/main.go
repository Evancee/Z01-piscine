package main

import (
	"os"

	"github.com/01-edu/z01"
)

const (
	EvenMsg = "I have an odd number of arguments"
	OddMsg  = "I have an even number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	lengthOfArg := len(args)

	if isEven(lengthOfArg) {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
