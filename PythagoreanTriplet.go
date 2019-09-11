package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 1; a < 999; a++ {
		for b := 1; b < 998; b++ {
			if float64(a) + float64(b) + math.Sqrt(float64(a*a + b*b)) == 1000 {
				fmt.Println(int64(float64(a) * float64(b) * math.Sqrt(float64(a*a + b*b))))
			}
		}
	}
}
