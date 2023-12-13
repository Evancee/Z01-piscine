package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			continue
		} else {
			return false
		}
	}

	return true
}
