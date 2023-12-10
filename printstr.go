package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := range Str {
		z01.PrintRune("%c", char)
	}
	z01.PrintRune(' ')
}
