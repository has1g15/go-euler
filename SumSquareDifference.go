package main

import (
	"fmt"
	"math"
)

func main() {
	x := 100
	fmt.Println(calculateSquareOfSums(x) - calculateSumOfSquares(x))
}

func calculateSumOfSquares(num int) int{
	total := 0
	for i := 1; i <= num; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return total
}

func calculateSquareOfSums(num int) int{
	total := 0
	for i := 1; i <= num; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2))
}