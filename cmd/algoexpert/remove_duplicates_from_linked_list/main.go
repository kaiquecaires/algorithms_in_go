package removeduplicatesfromlinkedlist

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (l *LinkedList) NextDiff() *LinkedList {
	firstValue := l.Value
	curr := l

	for curr != nil && curr.Value == firstValue {
		curr = curr.Next
	}

	return curr
}

func RemoveDuplicatesFromLinkedList(linkedList *LinkedList) *LinkedList {
	curr := linkedList

	for curr != nil {
		curr.Next = curr.NextDiff()
		curr = curr.Next
	}

	return linkedList
}
