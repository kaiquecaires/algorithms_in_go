package algoexpert

import "sort"

func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)
	changes := 0

	for _, coin := range coins {
		if coin <= changes+1 {
			changes += coin
		} else {
			break
		}
	}

	return changes + 1
}
