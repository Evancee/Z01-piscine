package piscine

import "github.com/01-edu/z01"

func main() {
	var p *int
	p = new(int)
	*p = 1
	z01.PrintRune(*p)
}
