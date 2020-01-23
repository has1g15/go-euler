package main

import (
	"fmt"
	"math"
)

func main() {
	x := [][]int32{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}
	matrixRotation(x, 40)
}

func matrixRotation(matrix [][]int32, r int32) {
	m := len(matrix)
	var n int
	for _, x := range matrix {
		n = len(x)
	}

	rings := (math.Min(float64(m), float64(n)))/2

	for x := 0; x < int(rings); x++ {
		for rot := 0; rot < int(r) % (2*(m-2*x) + 2*(n-2*x) -4); rot++ {
			topLeft := matrix[x][x]

			for i := x; i < n-(x+1); i++ {
				matrix[x][i] = matrix[x][i+1]
			}

			for i := x; i < m-(x+1); i++ {
				matrix[i][n-(x+1)] = matrix[i+1][n-(x+1)]
			}

			for i := n - (x+1); i > x; i-- {
				matrix[m-(x+1)][i] = matrix[m-(x+1)][i-1]
			}

			for i := m - (x+1); i > x; i-- {
				matrix[i][x] = matrix[i-1][x]
			}
			matrix[x+1][x] = topLeft
		}
	}

	for _, x := range matrix {
		for _, y := range x {
			fmt.Printf("%d ",y)
		}
		fmt.Println()
	}
}
