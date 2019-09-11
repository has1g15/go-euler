package main

import (
	"fmt"
	"math"
)

func main() {
	n := 10001
	x := 2
	count := 0

	for count < n {
		if isPrimeNum(x) != 0 {
			count++
		}
		x++
	}
	fmt.Println(isPrimeNum(x-1))
}

func isPrimeNum(x int) int {
	for y := int(math.Sqrt(float64(x))); y >= 1; y-- {
		if y == 1 {
			return x
		}
		if x % y == 0 {
			return 0
		}
	}
	return x
}
