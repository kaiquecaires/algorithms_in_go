package numberofwaystotraversegraph

func NumberOfWaysToTraverseGraph(width int, height int) int {
	graph := make([][]int, height)

	for j := 0; j < height; j++ {
		graph[j] = make([]int, width)
	}

	for j := 0; j < height; j++ {
		for i := 0; i < width; i++ {
			if i == 0 || j == 0 {
				graph[j][i] = 1
			} else {
				graph[j][i] = graph[j-1][i] + graph[j][i-1]
			}
		}
	}

	return graph[height-1][width-1]
}
