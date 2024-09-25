package reconstructbst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	bst := &BST{Value: preOrderTraversalValues[0]}

	for i := 1; i < len(preOrderTraversalValues); i++ {
		bst.Insert(preOrderTraversalValues[i])
	}

	return bst
}

func (tree *BST) Insert(value int) {
	currentNode := tree

	for currentNode != nil {
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
}
