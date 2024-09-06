package algoexpert

import (
	"math"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type ClosestNode struct {
	Value int
	Node  *BST
}

func (tree *BST) FindClosestValue(target int) int {
	currentNode := tree
	closestNode := &ClosestNode{
		Value: int(math.Abs(float64(target) - float64(currentNode.Value))),
		Node:  currentNode,
	}

	for currentNode != nil {
		value := int(math.Abs(float64(target) - float64(currentNode.Value)))
		if value < closestNode.Value {
			closestNode.Value = value
			closestNode.Node = currentNode
		}
		if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}

	return closestNode.Node.Value
}
