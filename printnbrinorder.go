package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n <= 0 {
		z01.PrintRune('0')
	} else {
		result := []int{}
		for n > 0 {
			result = append(result, n%10)
			n /= 10
		}
		for i := 0; i < len(result); i++ {
			for j := i + 1; j < len(result)-i-1; j++ {
				if result[j] > result[j+1] {
					result[j], result[j+1] = result[j+1], result[j]
				}
			}
		}
		for _, c := range result {
			z01.PrintRune(rune(c + '0'))
		}

	}
}
