package numberofwaystomakechange

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	ways := make([]int, n+1)
	ways[0] = 1

	for _, denom := range denoms {
		for i := range ways {
			if denom <= i {
				ways[i] += ways[i-denom]
			}
		}
	}

	return ways[n]
}
