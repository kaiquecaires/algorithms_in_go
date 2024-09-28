package findsuccessor

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	if node.Right != nil {
		currentNode := node.Right
		for currentNode.Left != nil {
			currentNode = currentNode.Left
		}
		return currentNode
	}

	currentNode := node.Parent
	for currentNode != nil && currentNode.Right == node {
		currentNode = currentNode.Parent
	}

	return currentNode
}
