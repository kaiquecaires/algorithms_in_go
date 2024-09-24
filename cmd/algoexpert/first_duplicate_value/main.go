package firstduplicatevalue

import (
	"math"
)

func FirstDuplicateValue(array []int) int {
	seenIdx := 0

	for i := 0; i < len(array); i++ {
		seenIdx = int(math.Abs(float64(array[i])) - 1)

		if array[seenIdx] < 0 {
			return int(math.Abs(float64(array[i])))
		}

		array[seenIdx] = 0 - array[seenIdx]
	}

	return -1
}
