package evaluateexpressiontree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func EvaluateExpressionTree(tree *BinaryTree) int {
	result := tree.Value

	switch result {
	case -1:
		return EvaluateExpressionTree(tree.Left) + EvaluateExpressionTree(tree.Right)
	case -2:
		return EvaluateExpressionTree(tree.Left) - EvaluateExpressionTree(tree.Right)
	case -3:
		return EvaluateExpressionTree(tree.Left) / EvaluateExpressionTree(tree.Right)
	case -4:
		return EvaluateExpressionTree(tree.Left) * EvaluateExpressionTree(tree.Right)
	default:
		return result
	}
}
