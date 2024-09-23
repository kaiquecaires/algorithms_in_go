package spiraltraverse

func SpiralTraverse(array [][]int) []int {
	output := []int{}

	left, right := 0, len(array[0])-1
	top, bottom := 0, len(array)-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			output = append(output, array[top][i])
		}

		if top == bottom {
			break
		}

		top = top + 1

		for i := top; i <= bottom; i++ {
			output = append(output, array[i][right])
		}

		if left == right {
			break
		}

		right = right - 1

		for i := right; i >= left; i-- {
			output = append(output, array[bottom][i])
		}

		bottom = bottom - 1

		for i := bottom; i >= top; i-- {
			output = append(output, array[i][left])
		}

		left = left + 1
	}

	return output
}
