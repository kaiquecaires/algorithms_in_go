package getnthfib

func GetNthFib(n int) int {
	lastTwo := [2]int{0, 1}

	if n == 1 {
		return lastTwo[0]
	}

	if n == 2 {
		return lastTwo[1]
	}

	for i := 2; i < n; i++ {
		lastTwo[0], lastTwo[1] = lastTwo[1], lastTwo[0]+lastTwo[1]
	}

	return lastTwo[1]
}
