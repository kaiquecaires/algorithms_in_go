package middlenode

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func MiddleNode(linkedList *LinkedList) *LinkedList {
	fast, slow := linkedList, linkedList

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
