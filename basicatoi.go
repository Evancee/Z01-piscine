package piscine

func BasicAtoi(s string) int {
	var i int
	for _, char := range s {
		i = i*10 + int(char-'0')
	}
	return i
}
