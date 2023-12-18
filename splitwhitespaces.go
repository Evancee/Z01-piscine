package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	begin := 0
	for i, r := range s {
		if separator(r) {
			if i > begin {
				words = append(words, s[begin:i])
			}
			begin = i + 1
		}
	}
	if begin < len(s) {
		words = append(words, s[begin:])
	}
	return words
}

func separator(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n'
}
