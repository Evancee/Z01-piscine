package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
