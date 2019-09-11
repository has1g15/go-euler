package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	total := 2
	start := time.Now()

	jobs := make(chan int, 2000000)
	results := make(chan int, 2000000)

	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)
	go worker(jobs, results)

	for i := 3; i <= 2000000; i+=2 {
		jobs <- i
	}
	close(jobs)

	for j := 3; j <= 2000000; j+=2 {
		total += <- results
	}

	fmt.Println(total)
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func worker(jobs <- chan int, results chan <- int) {
	for num := range jobs {
		results <- isPrimeNum(num)
	}
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