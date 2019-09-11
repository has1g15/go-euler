package main

import (
	"fmt"
	"math"
)

func main() {
	num := 1
	count := 1
	for findDivisorNumber(num) < 500{
		count++
		num += count
	}
	fmt.Printf("Triangle number: %d\n", num)
	fmt.Printf("Number of divisors: %d\n", findDivisorNumber(num))
}

func findDivisorNumber(num int) int {
	divisorNum := 0
	for i:=1; i <= int(math.Sqrt(float64(num))); i++ {
		if num % i == 0 {
			divisorNum += 2
			if math.Pow(float64(i),2) == float64(num) {
				divisorNum--
			}
		}
	}
	return divisorNum
}