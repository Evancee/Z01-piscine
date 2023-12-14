package piscine

func CountIf(f func(string) bool, tab []string) int {
	counti := 0
	for _, tab := range tab {
		if f(tab) {
			counti++
		}
	}
	return counti
}
