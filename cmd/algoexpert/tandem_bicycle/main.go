package tandembicycle

import (
	"math"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Ints(blueShirtSpeeds)

	if fastest {
		return max(redShirtSpeeds, blueShirtSpeeds)
	} else {
		return min(redShirtSpeeds, blueShirtSpeeds)
	}
}

func max(redShirtSpeeds []int, blueShirtSpeeds []int) int {
	sum := 0
	for i, j := 0, len(redShirtSpeeds)-1; i < len(redShirtSpeeds); i, j = i+1, j-1 {
		sum += int(math.Max(float64(redShirtSpeeds[i]), float64(blueShirtSpeeds[j])))
	}
	return sum
}

func min(redShirtSpeeds []int, blueShirtSpeeds []int) int {
	sum := 0
	for i := 0; i < len(redShirtSpeeds); i++ {
		sum += int(math.Max(float64(redShirtSpeeds[i]), float64(blueShirtSpeeds[i])))
	}
	return sum
}
