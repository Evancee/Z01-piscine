package piscine

func TrimAtoi(s string) int {
	if len(s) == 0 || s[0] == '-' && len(s) == 1 {
		return 0
	}
	isNegative := false
	if s[0] == '-' {
		isNegative = true
		s = s[1:]
	}
	num := 0
	for _, v := range s {
		if v < '0' || v > '9' {
			break
		}
		num = num*10 + int(v) - 48
	}
	if isNegative {
		num *= -1
	}
	return num
}
