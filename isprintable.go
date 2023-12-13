package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if char >= ' ' && char <= '~' {
			continue
		} else {
			return false
		}
	}
	return true
}
