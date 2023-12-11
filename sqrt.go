package piscine

func Sqrt(nb int) int {
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return 1
		}
	}
	return 0
}
