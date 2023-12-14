package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	var a, b, operator string

	a = os.Args[1]
	operator = os.Args[2]
	b = os.Args[3]
	var valueA, valueB float64
	var err error

	valueA, err = strconv.ParseFloat(a, 64)
	if err != nil {
		return
	}

	valueB, err = strconv.ParseFloat(b, 64)
	if err != nil {
		return
	}
	var result float64
	switch operator {
	case "+":
		result = valueA + valueB
	case "-":
		result = valueA - valueB
	case "*":
		result = valueA * valueB
	case "/":
		if valueB == 0 {
			return
		}
		result = valueA / valueB
	case "%":
		if valueB == 0 {
			return
		}
		result = valueB
	default:
		return
	}

	if result == math.Inf(1) || result == math.Inf(-1) {
		return
	}

	fmt.Printf("%g\n", result)
}
