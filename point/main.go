package main

import (
	"strconv"

	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')

	xStr := strconv.Itoa(points.x)
	yStr := strconv.Itoa(points.y)

	for _, r := range xStr {
		z01.PrintRune(r)
	}
	z01.PrintRune(',')
	z01.PrintRune(' ')
	for _, r := range yStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
