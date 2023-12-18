package piscine

func ConcatParams(args []string) string {
	var fstring string
	for index, item := range args {
		fstring += string(item)
		if index != len(args)-1 {
			fstring += " \n"
		}
	}
	return fstring
}
