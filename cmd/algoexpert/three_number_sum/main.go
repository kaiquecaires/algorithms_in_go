package threenumbersum

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	result := [][]int{}
	sort.Ints(array)

	for i := 0; i < len(array); i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum > target {
				right--
			} else if currentSum < target {
				left++
			} else {
				result = append(result, []int{array[i], array[left], array[right]})
				left++
			}
		}
	}

	return result
}
