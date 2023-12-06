package piscine

func Sqrt(nb int) int {
	x := nb
	y := 1

	for (x - y) != 0 {
		x = (x + y) / 2
		y = nb / x
	}
	return x
}
