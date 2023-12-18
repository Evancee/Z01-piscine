package piscine

func ConcatParams(args []string) string {
	var answer string
	for i := range args {
		answer += args[1]
		if i != len(args)-1 {
			answer += " \n"
		}
	}
	return answer
}
