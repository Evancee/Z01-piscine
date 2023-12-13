package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	tbl := []int(nil)
	for n > 0 {
		tbl = append(tbl, n%10)
		n = n / 10
	}
	for swapp := true; swapp; {
		swapp := false
		for i := 1; i < len(tbl); i++ {
			if tbl[i-1] > tbl[i] {
				tbl[i-1], tbl[i] = tbl[i], tbl[i-1]
				swapp = true
			}
		}
	}
	for _, v := range tbl {
		z01.PrintRune(rune(v + '0'))
	}
}
