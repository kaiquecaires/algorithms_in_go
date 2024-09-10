package findthreelargestnumbers

import "math"

func FindThreeLargestNumbers(array []int) []int {
	threeLargestNumbers := []int{math.MinInt, math.MinInt, math.MinInt}

	for _, num := range array {
		var prev int = num
		for i := 2; i >= 0; i-- {
			threeLargestNumbers[i], prev =
				int(math.Max(float64(threeLargestNumbers[i]), float64(prev))),
				int(math.Min(float64(threeLargestNumbers[i]), float64(prev)))
		}
	}

	return threeLargestNumbers
}
