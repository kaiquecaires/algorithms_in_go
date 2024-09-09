package productsum

type SpecialArray []interface{}

// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return Sum(array, 1)
}

func Sum(array []interface{}, depth int) int {
	sum := 0

	for _, n := range array {
		switch casted := n.(type) {
		case int:
			sum += casted
		case SpecialArray:
			sum += Sum(casted, depth+1)
		}
	}

	return sum * depth
}
