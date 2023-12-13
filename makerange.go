package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var result []int
	result = make([]int, min-max)
	for i := 0; i < max-min; i++ {
		result[i] = i + min
	}
	return result
}
