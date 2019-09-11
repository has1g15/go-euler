package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	x := 600851475143

	for y := int(math.Sqrt(float64(x))); y >= 1; y-- {
		if x % y == 0 && isPrime(y) {
			if digitCount(y) > 3 {
				fmt.Println(formatNumber(y, digitCount(y)))
			} else {
				fmt.Println(y)
			}
			break
		}
	}
}

func isPrime(x int) bool {
	for y := int(math.Sqrt(float64(x))); y >= 1; y-- {
		if y == 1 {
			return true
		}
		if x % y == 0 {
			return false
		}
	}
	return true
}

func formatNumber(x int, count int) string {
	numToFormat := strconv.Itoa(x)

	for count > 3 {
		count -= 3
		numToFormat = numToFormat[:count] + "," + numToFormat[count:]
	}
	return numToFormat
}

func digitCount(i int) int {
	count := 0
	for i != 0 {
		i /= 10
		count += 1
	}
	return count
}
