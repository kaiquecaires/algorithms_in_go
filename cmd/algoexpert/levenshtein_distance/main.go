package levenshteindistance

import "strings"

func LevenshteinDistance(a, b string) int {
	cache := make([][]int, len(a)+1)

	for i := range cache {
		cache[i] = make([]int, len(b)+1)

		for j := range cache[i] {
			cache[i][j] = 1 << 62
		}
	}

	// fill the base cases
	for j := 0; j < len(b)+1; j++ {
		cache[len(a)][j] = len(b) - j
	}

	for i := 0; i < len(a)+1; i++ {
		cache[i][len(b)] = len(a) - i
	}

	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")

	for i := len(a) - 1; i >= 0; i-- {
		for j := len(b) - 1; j >= 0; j-- {
			if aArr[i] == bArr[j] {
				cache[i][j] = cache[i+1][j+1]
			} else {
				cache[i][j] = 1 + Min(cache[i+1][j], cache[i][j+1], cache[i+1][j+1])
			}
		}
	}

	return cache[0][0]
}

func Min(nums ...int) int {
	min := nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
	}
	return min
}
