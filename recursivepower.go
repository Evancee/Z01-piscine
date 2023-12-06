package piscine

func RecursivePower(nb int, power int) int {
	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	if power%2 == 0 {
		return RecursivePower(nb*nb, power/2)
	}

	return nb * RecursivePower(nb, power-1)
}
