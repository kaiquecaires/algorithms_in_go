package kadanesalgorithm

import (
	"math"
)

func KadanesAlgorithm(array []int) int {
	maxEndingHere := array[0]
	maxSoFar := array[0]

	for i := 1; i < len(array); i++ {
		currentSum := array[i] + maxEndingHere
		maxEndingHere = int(math.Max(float64(currentSum), float64(array[i])))
		maxSoFar = int(math.Max(float64(maxSoFar), float64(maxEndingHere)))
	}

	return maxSoFar
}
