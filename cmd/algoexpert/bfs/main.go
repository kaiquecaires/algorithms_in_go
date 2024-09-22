package bfs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	queue := []*Node{n}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		array, queue = append(array, node.Name), append(queue, node.Children...)
	}

	return array
}
