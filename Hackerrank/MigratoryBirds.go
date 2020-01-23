package main

import "fmt"

func main() {
	y := []int32 {100,100,50,40,40,20,10}
	fmt.Println(migratoryBirds(y))
}

func migratoryBirds(arr []int32) int32 {
	birdCount := make(map[int32]int)

	for i, _ := range arr {
		if _, ok := birdCount[arr[i]]; ok {
			birdCount[arr[i]]++
		} else {
			birdCount[arr[i]] = 1
		}
	}

	maxFrequency := 0
	for _, v := range birdCount {
		if v > maxFrequency {
			maxFrequency = v
		}
	}

	var maxIndex int32 = 0
	for _, v := range arr {
		if v > maxIndex {
			maxIndex = v
		}
	}

	minIndex := maxIndex
	for k, v := range birdCount {
		if maxFrequency == v {
			if k < minIndex {
				minIndex = k
			}
		}
	}

	return minIndex
}