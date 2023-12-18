package piscine

func MakeRange(min, max int) []int {
	result := max - min
	if result <= 0 {
		return nil
	}
	ans := make([]int, result)
	for i := range ans {
		ans[i] = min
		min++
	}
	return ans
}
