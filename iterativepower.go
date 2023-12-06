package piscine

func IterativePower(nb int, power int) int {
	result := 0
	if power < 0 {
		return 0
	}
	for i := 1; i < power; i++ {
		result *= nb
	}
	return result
}
