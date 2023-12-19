package piscine

func Capitalize(s string) string {
	s = ToLower(s)
	for i, v := range s {
		if i == 0 {
			s = ToUpper(string(v)) + s[1+1:]
		} else {
			if isAlpha(string(v)) && !IsAlpha(string(s[i-1])) {
				if i != len(s)-1 {
					s = s[:i] + ToUpper(string(v)) + s[i+1:]
				} else {
					s = s[:i] + ToUpper(string(v))
				}
			}
		}
	}
}
