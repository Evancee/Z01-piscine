package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var sorted bool
	sorted = true
	for i := 0; i < len(tab)-1; i++{
		if f(tab[i], tab[i+1]) >0{
			sorted = false
		}
	}
}
if sorted == false {
	sorted = true
	for i := 0; i< len(tab)-1; i++{
		if f(tab[i], tab[i+1])<0{
			sorted = false
		}
	}
	return sorted
}