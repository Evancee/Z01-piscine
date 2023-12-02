package main

import "fmt"

func giveOne(n *int) {
	*n = 1
}

func main() {
	myNum := 5
	fmt.Println("Before: ", myNum)
	giveOne(&myNum)
	fmt.Println("After: ", myNum)
}
