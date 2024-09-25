package validatebst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) ValidateBst() bool {
	min, max := -(1 << 62), 1<<62
	return validateHelper(tree, min, max)
}

func validateHelper(tree *BST, min int, max int) bool {
	if tree == nil {
		return true
	}

	if tree.Value < min || tree.Value >= max {
		return false
	}

	return validateHelper(tree.Left, min, tree.Value) && validateHelper(tree.Right, tree.Value, max)
}
