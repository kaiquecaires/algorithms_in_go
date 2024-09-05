package algoexpert

import "math"

func SortedSquaredArray(array []int) []int {
	smallest := 0
	largest := len(array) - 1
	var output = make([]int, len(array))
	lastIdx := len(array) - 1

	for lastIdx >= 0 {
		if math.Abs(float64(array[smallest])) > math.Abs(float64(array[largest])) {
			output[lastIdx] = array[smallest] * array[smallest]
			smallest++
		} else {
			output[lastIdx] = array[largest] * array[largest]
			largest--
		}
		lastIdx--
	}

	return output
}
