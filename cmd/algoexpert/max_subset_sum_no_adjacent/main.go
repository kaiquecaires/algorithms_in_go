package maxsubsetsumnoadjacent

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}

	if len(array) == 1 {
		return array[0]
	}

	first, second := array[0], Max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		first, second = second, Max(second, array[i]+first)
	}

	return second
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
