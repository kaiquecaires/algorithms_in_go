package ismonotonicarray

import (
	"math"
)

func IsMonotonic(array []int) bool {
	if len(array) == 0 || len(array) == 1 {
		return true
	}

	increasing := array[len(array)-1] > array[0]
	ref := array[0]

	for i := 0; i < len(array); i++ {
		if increasing {
			if array[i] < ref {
				return false
			}
			ref = int(math.Max(float64(ref), float64(array[i])))
		} else {
			if array[i] > ref {
				return false
			}
			ref = int(math.Min(float64(ref), float64(array[i])))
		}
	}

	return true
}
