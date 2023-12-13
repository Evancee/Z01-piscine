package piscine

func AlphaCount(s string) int {
	var acount int
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			acount++
		}
	}
	return acount
}
