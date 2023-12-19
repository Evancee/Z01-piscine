package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	for i, v := range s {
		if v >= 'A' && v <= 'z' {
			runes[i] = v + 32
		}
	}
	for index, value := range runes {
		if value >= 'a' && value <= 'z' && index != 0 {
			if (runes[index-1] < 'A' || runes[index-1] > 'Z') && (runes[index-1] < 'a' || runes[index-1] > 'z') && (runes[index-1] < '0' || runes[index-1] > '9') {
				runes[index] = value - 32
			}
		} else if index == 0 && value >= 'a' && value <= 'z' {
			runes[index] = value - 32
		}
	}
	return string(runes)
}
