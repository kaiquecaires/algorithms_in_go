package minheightbst

func MinHeightBST(array []int) *BST {
	return InsertHelper(nil, array)
}

func InsertHelper(bst *BST, array []int) *BST {
	if len(array) == 0 {
		return bst
	}

	mid := len(array) / 2

	if bst == nil {
		bst = &BST{Value: array[mid]}
	} else {
		bst.Insert(array[mid])
	}

	InsertHelper(bst, array[:mid])
	InsertHelper(bst, array[mid+1:])
	return bst
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
