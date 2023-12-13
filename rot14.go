package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		if char >= 'A' && char <= 'Z' {
			if char+14 > 'Z' {
				result += string(char + 14 - 26)
			} else {
				result += string(char + 14)
			}
		} else if char >= 'a' && char <= 'z' {
			if char+14 > 'z' {
				result += string(char + 14 - 26)
			} else {
				result += string(char + 14)
			}
		} else {
			result += string(char)
		}
	}
	return result
}
