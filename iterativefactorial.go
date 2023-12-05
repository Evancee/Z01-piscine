package piscine

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune(IterativeFactorial(6))
}

func IterativeFactorial(nb int) int {
	result := 1
	if nb == 0 || nb == 1 {
		return 1
	}
	for i := 1; i <= nb; i++ {
		result = result * i
	}
	return result
}
