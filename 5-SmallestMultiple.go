package main

import "fmt"

func main() {
	num := 2520
	numFound := false
	for !numFound {
		num+=20
		numFound = isDivisibleFor20(num)
	}
	fmt.Println(num)
}

func isDivisibleFor20(num int) bool {
	dividing := true
	for i:= 2; i <= 20; i++ {
		if num % i != 0 {
			dividing = false
			break
		}
	}
	return dividing
}