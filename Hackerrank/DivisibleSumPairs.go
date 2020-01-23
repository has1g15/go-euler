package main

import "fmt"

func main() {
	y := []int32 {100,100,50,40,40,20,10}
	fmt.Println(divisibleSumPairs(6,3, y))
}

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var pairs int32
	for i:= 0; i < int(n) - 1; i++ {
		for j := i+1; j < int(n); j++ {
			if (ar[i] + ar[j]) % k == 0 {
				pairs++
			}
		}
	}
	return pairs
}
