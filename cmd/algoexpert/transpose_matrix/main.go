package transposematrix

func TransposeMatrix(matrix [][]int) [][]int {
	transposedMatrix := make([][]int, len(matrix[0]))

	for _, rowValue := range matrix {
		for col, colValue := range rowValue {
			transposedMatrix[col] = append(transposedMatrix[col], colValue)
		}
	}

	return transposedMatrix
}
