package piscine

func MakeRange(min, max int) []int {
	if max > min {
		result := make([]int, max-min)
		for i := min; i < max-min; i++ {
			result[i] += min + i
		}
		return result
	} else {
		return nil
	}
}
