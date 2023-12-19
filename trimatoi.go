package piscine

func TrimAtoi(s string) int {
	isNegative := false
	number := 0

	for _, char := range s {
		if char == '-' {
			isNegative = true
		} else if char >= '0' && char <= '9' {
			number = number*10 + int(char-'0')
		} else {
			return 0
		}
	}

	if isNegative {
		number = -number
	}

	return number
}
