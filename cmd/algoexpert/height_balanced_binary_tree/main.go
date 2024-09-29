package heightbalancedbinarytree

import "math"

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	return Helper(tree) != -1
}

func Helper(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	left := Helper(tree.Left)
	right := Helper(tree.Right)

	if int(math.Abs(float64(left-right))) > 1 {
		return -1
	}

	max := int(math.Max(float64(left), float64(right)))

	return max + 1
}
