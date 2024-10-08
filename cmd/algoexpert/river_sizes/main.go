package riversizes

func RiverSizes(matrix [][]int) []int {
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	result := []int{}

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 1 {
				result = append(result, dfs(matrix, row, col, directions))
			}
		}
	}

	return result
}

func dfs(matrix [][]int, row, col int, directions [][2]int) int {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] == 0 {
		return 0
	}

	matrix[row][col] = 0
	size := 1

	for _, dir := range directions {
		size += dfs(matrix, row+dir[0], col+dir[1], directions)
	}

	return size
}
