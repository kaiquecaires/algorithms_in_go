package longestpeak

func LongestPeak(array []int) int {
	longestPeak, i := 0, 1

	for i < len(array)-1 {
		isPeak := array[i] > array[i-1] && array[i] > array[i+1]
		if !isPeak {
			i++
			continue
		}

		left := i - 1
		for left > 0 {
			if array[left] <= array[left-1] {
				break
			}
			left--
		}

		right := i + 1
		for right < len(array)-1 {
			if array[right] <= array[right+1] {
				break
			}
			right++
		}

		i = right
		longestPeak = Max(longestPeak, right-left+1)
	}

	return longestPeak
}

func Max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}
