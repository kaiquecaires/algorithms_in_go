package reverselinkedlist

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func ReverseLinkedList(head *LinkedList) *LinkedList {
	currentNode := head
	var prev *LinkedList

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	return prev
}
