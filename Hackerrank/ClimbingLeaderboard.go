package main

import "fmt"

func main() {
	y := []int32 {100,100,50,40,40,20,10}
	alice := []int32 {5, 25, 55, 102}
	fmt.Println(climbingLeaderboard(y, alice))
}

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	currentRank := 1
	ranks := make([]int, 0, len(scores))
	aliceRanks := make([]int32, 0, len(alice))

	ranks = append(ranks, 1)
	for i := 1; i < len(scores); i++ {
		if !(scores[i] == scores[i-1]) {
			currentRank++
		}
		ranks = append(ranks, currentRank)
	}

	i := len(scores) - 1
	for _, v := range alice {
		for j := i; j >= 0; j-- {
			i = j
			if v < scores[j] {
				aliceRanks = append(aliceRanks, int32(ranks[j] + 1))
				break
			} else if v == scores[j] {
				aliceRanks = append(aliceRanks, int32(ranks[j]))
				break
			} else if v > scores[0] {
				aliceRanks = append(aliceRanks, 1)
				break
			}
		}
	}

	return aliceRanks
}
