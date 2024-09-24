package arrayofproducts

func ArrayOfProducts(array []int) []int {
	right := make([]int, len(array))
	left := make([]int, len(array))
	products := make([]int, len(array))

	for i := 0; i < len(array); i++ {
		right[i], left[i], products[i] = 1, 1, 1
	}

	product := 1
	for i := 0; i < len(array); i++ {
		right[i] = product
		product *= array[i]
	}

	product = 1
	for i := len(array) - 1; i >= 0; i-- {
		left[i] = product
		product *= array[i]
	}

	for i := 0; i < len(array); i++ {
		products[i] = left[i] * right[i]
	}

	return products
}
