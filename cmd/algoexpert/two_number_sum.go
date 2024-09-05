package algoexpert

func TwoNumberSum(array []int, target int) []int {
	expectedNumbers := map[int]int{}

	for i, num := range array {
		expectedNumber := target - num
		if _, exists := expectedNumbers[expectedNumber]; exists {
			return []int{num, expectedNumber}
		}
		expectedNumbers[num] = i
	}

	return []int{}
}
