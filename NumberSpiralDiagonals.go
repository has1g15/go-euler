package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()
	gridWidth := 1001
	generateDiagonalSums((gridWidth + 1)/2)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func generateDiagonalSums(x int) {
	sums := []int{1}

	for i := 1; i <= x; i++ {
		sums = append(sums, sumDiagonals(i, sums))
	}
	fmt.Println(sums[x-1])
}

func sumDiagonals(x int, sums []int) int {
	return (16 * int(math.Pow(float64(x), 2))) + (4*x) + 4 + sums[x-1]
}