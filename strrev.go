package piscine

func StrRev(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i++ {
		reversed += string(s[i])
	}
	return reversed
}
