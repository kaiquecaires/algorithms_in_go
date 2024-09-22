package smallestdifference

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	smallest := math.Inf(1)
	result := make([]int, 2)

	p1, p2 := 0, 0

	for p1 < len(array1) && p2 < len(array2) {
		absoluteValue := math.Abs(float64(array1[p1] - array2[p2]))

		if absoluteValue < smallest {
			smallest = absoluteValue
			result[0] = array1[p1]
			result[1] = array2[p2]
		}

		if array1[p1] < array2[p2] && p1 < len(array1) {
			p1++
		} else {
			p2++
		}
	}

	return result
}
