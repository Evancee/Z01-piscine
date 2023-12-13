package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, min-max)
	for i := min; i < max; i++ {
		result[i-min] = i
	}
	return result
}
