package binarytreediameter

import "math"

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func BinaryTreeDiameter(tree *BinaryTree) int {
	var max int
	Helper(tree, 0, &max)
	return max
}

func Helper(tree *BinaryTree, d int, max *int) int {
	if tree == nil {
		return 0
	}

	left := Helper(tree.Left, d+1, max)
	right := Helper(tree.Right, d+1, max)
	maxToReturn := int(math.Max(float64(left), float64(right)))
	maxToReturn = int(math.Max(float64(maxToReturn), float64(d)))

	*max = int(math.Max(float64(*max), float64((left-d)+(right-d))))

	return maxToReturn
}
