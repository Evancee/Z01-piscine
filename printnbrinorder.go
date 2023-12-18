package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	for n > 0 {

		digit := n % 10

		runeDigit := rune(digit)

		z01.PrintRune(runeDigit + '1')

		n /= 10
	}
}
