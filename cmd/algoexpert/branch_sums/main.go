package branchsums

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var array []int
	return BranchSumsHelper(root, 0, array)
}

func BranchSumsHelper(root *BinaryTree, sum int, array []int) []int {
	sum += root.Value

	if root.Left == nil && root.Right == nil {
		array = append(array, sum)
		return array
	}

	if root.Left != nil {
		array = BranchSumsHelper(root.Left, sum, array)
	}

	if root.Right != nil {
		array = BranchSumsHelper(root.Right, sum, array)
	}

	return array
}
