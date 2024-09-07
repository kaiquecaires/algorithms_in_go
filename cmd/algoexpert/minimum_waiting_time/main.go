package minimumwaitingtime

import (
	"sort"
)

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	waitingTime, minWaitingTime := 0, 0

	for _, query := range queries {
		minWaitingTime += waitingTime
		waitingTime += query
	}

	return minWaitingTime
}
