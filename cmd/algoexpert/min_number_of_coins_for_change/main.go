package minnumberofcoinsforchange

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	changes := make([]int, n+1)

	for i := range changes {
		changes[i] = 1 << 62
	}
	changes[0] = 0

	for _, denom := range denoms {
		for i := denom; i < len(changes); i++ {
			changes[i] = int(math.Min(float64(1+changes[(i-denom)]), float64(changes[i])))
		}
	}

	if changes[n] == 1<<62 {
		return -1
	} else {
		return changes[n]
	}
}
