package piscine

func TrimAtoi(s string) int {
	index := 0
	sign := 1
	for _, v := range s {
		if v >= '0' && v <= '9' {
			v = v - '0'
			index = index*10 + int(v)
		} else if index == 0 && v == '-' {
			sign = -1
		}
	}
	return sign * index
}
