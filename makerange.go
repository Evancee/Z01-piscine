package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	result := make([]int, max-min+1)
	for i := max - min; i >= 0; i-- {
		result[i] = max - 1
	}
	return result
}
