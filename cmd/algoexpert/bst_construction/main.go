package bstconstruction

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	currentNode := tree
	for {
		if value >= currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = &BST{Value: value}
				break
			}
			currentNode = currentNode.Right
		} else {
			if currentNode.Left == nil {
				currentNode.Left = &BST{Value: value}
				break
			}
			currentNode = currentNode.Left
		}
	}

	return tree
}

func (tree *BST) Contains(value int) bool {
	currentNode := tree

	for currentNode != nil {
		if value > currentNode.Value {
			currentNode = currentNode.Right
		} else if value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			return true
		}
	}

	return false
}

func (tree *BST) Remove(value int) *BST {
	if tree == nil {
		return nil
	}

	if value > tree.Value {
		tree.Right = tree.Right.Remove(value)
		return tree
	} else if value < tree.Value {
		tree.Left = tree.Left.Remove(value)
		return tree
	} else {
		if tree.Right == nil && tree.Left == nil {
			return nil
		}

		if tree.Right == nil {
			*tree = *tree.Left
			return tree
		}

		if tree.Left == nil {
			*tree = *tree.Right
			return tree
		}
	}

	currentNode := tree.Right

	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	tree.Value = currentNode.Value
	tree.Right = tree.Right.Remove(tree.Value)
	return tree
}
