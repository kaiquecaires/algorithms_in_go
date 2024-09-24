package mergeoverlapintervals

import (
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	min, max := intervals[0][0], intervals[0][1]
	output := [][]int{}

	for i := 1; i < len(intervals); i++ {
		cMin, cMax := intervals[i][0], intervals[i][1]

		if max >= cMin {
			max = Max(max, cMax)
		} else {
			output = append(output, []int{min, max})
			min, max = cMin, cMax
		}
	}

	output = append(output, []int{min, max})

	return output
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
