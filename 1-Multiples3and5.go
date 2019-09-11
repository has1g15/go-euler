package main

import "fmt"

func main() {
	fmt.Print(sumMultiples(1000))
}

func sumMultiples(num int) int{
	total := 0

	for i:=1; i < num; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	return total
}
