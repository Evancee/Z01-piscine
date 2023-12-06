package piscine

func powerRecursive(nb int, power int) int {
	if power == 0 {
		return 1
	}

	if power < 0 {
		return 0
	}

	if power%2 == 0 {
		return powerRecursive(nb*nb, power/2)
	}

	return nb * powerRecursive(nb, power-1)
}
