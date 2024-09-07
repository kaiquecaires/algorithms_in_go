package nodedepth

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func SumDepths(node *BinaryTree, depth int) int {
	sum := depth

	if node.Left != nil {
		sum += SumDepths(node.Left, depth+1)
	}

	if node.Right != nil {
		sum += SumDepths(node.Right, depth+1)
	}

	return sum
}

func NodeDepths(root *BinaryTree) int {
	output := SumDepths(root, 0)
	return output
}
