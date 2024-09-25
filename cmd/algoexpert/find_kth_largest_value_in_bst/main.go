package findkthlargestvalueinbst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func FindKthLargestValueInBst(tree *BST, k int) int {
	values := TraverseBST(tree, []int{})
	return values[k-1]
}

func TraverseBST(tree *BST, array []int) []int {
	if tree == nil {
		return array
	}

	array = TraverseBST(tree.Right, array)
	array = append(array, tree.Value)
	array = TraverseBST(tree.Left, array)
	return array
}
