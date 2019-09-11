package main

import (
	"fmt"
)

func main() {
	fmt.Println(calculateEvenFibonacci())
}

func calculateEvenFibonacci() int{
	total := 0
	sequence := []int {1, 2}
	max := 4000000
	i := 0
	nextTerm := 0

	for nextTerm < max {
		nextTerm = sequence[i]+sequence[i+1]
		if nextTerm % 2 == 0 {
			total += nextTerm
		}
		sequence = append(sequence, sequence[i]+sequence[i+1])
		i++
	}
	return total + 2
}
