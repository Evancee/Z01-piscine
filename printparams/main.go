package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	for _, str := range args {
		for _, char := range str {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
