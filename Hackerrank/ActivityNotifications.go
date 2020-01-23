package main

import (
	"fmt"
	sort2 "sort"
)

func main() {
	x := []int32{12, 3, 4, 2, 3, 6, 8, 4, 5}
	fmt.Println(activityNotifications(x, 5))
}

func activityNotifications(expenditure []int32, d int32) int32 {
	var fraudCount int32 = 0
	transactionData := make([]int32, d)
	copy(transactionData, expenditure[0:d])

	sort2.Slice(transactionData, func(i, j int) bool { return transactionData[i] < transactionData[j] })

	if isFraud(transactionData, expenditure[d]) {
		fraudCount++
	}

	for i := d; i < int32(len(expenditure)) - 1; i++ {
		itemToRemove := expenditure[i-d]
		itemToAdd := expenditure[i]
		if !(itemToAdd == itemToRemove) {
			transactionData = sortData(transactionData, itemToRemove, itemToAdd)
		}
		if isFraud(transactionData, expenditure[i+1]) {
			fraudCount++
		}
	}
	return fraudCount
}

func sortData(transactionData []int32, itemToRemove int32, itemToAdd int32) []int32 {
	x := sort2.Search(len(transactionData), func(i int) bool { return transactionData[i] > itemToRemove })
	transactionData = append(transactionData[:x-1], transactionData[x:]...)
	return insertSorted(transactionData, itemToAdd)
}

func isFraud(transactionData []int32, currentTransaction int32) bool {
	if (2 * calcMedian(transactionData)) <= float32(currentTransaction) {
		return true
	} else {
		return false
	}
}

func calcMedian(transactionData []int32) float32 {
	n := len(transactionData)
	if n % 2 == 0 {
		return float32((float32(transactionData[n/2-1]) + float32(transactionData[n/2])) / 2)
	} else {
		return float32(transactionData[n/2])
	}
}

func insertSorted(s []int32, e int32) []int32 {
	i := sort2.Search(len(s), func(i int) bool { return s[i] > e })
	s = append(s, 0)
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}