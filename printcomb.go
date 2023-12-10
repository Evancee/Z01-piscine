package piscine

import "github.com/01-edu/z01"

func printcomb() {
	for i := 0; i <= 7; i++ {
		for j := i + 1; j <= 8; j++ {
			for k := j + 1; k <= 9; k++ {
				z01.PrintRune("%d%d%d, ", i, j, k)
			}
		}
	}
	z01.PrintRune(' ')
}
