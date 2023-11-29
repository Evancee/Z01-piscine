package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums []int
	for i := 100; i < 1000; i++ {
		if isValid(i) {
			nums = append(nums, i)
		}
	}
	sort.Ints(nums)
	printCombos(nums)
}

func isValid(num int) bool {
	digits := getDigits(num)
	return digits[0] < digits[1] && digits[1] < digits[2]
}

func getDigits(num int) []int {
	return []int{num / 100, (num / 10) % 10, num % 10}
}

func printCombos(nums []int) {
	for i, num := range nums {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(num)
	}
	fmt.Println()
}
