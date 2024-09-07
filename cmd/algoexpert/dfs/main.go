package dfs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	array = append(array, n.Name)
	for _, node := range n.Children {
		array = node.DepthFirstSearch(array)
	}
	return array
}
