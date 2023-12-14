package piscine

func Any(f func(string) bool, a []string) bool {
	for _, a := range a {
		if f(a) {
			return true
		}
	}
	return false
}
