package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

// newNode creates a new tree node
func newNode(val int) *tree {
	return &tree{
		val:   val,
		left:  nil,
		right: nil,
	}
}

// buildTreeFromArray builds a binary tree from level-order array
func buildTreeFromArray(arr []interface{}) *tree {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}

	root := newNode(arr[0].(int))
	queue := []*tree{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// Process left child
		if i < len(arr) && arr[i] != nil {
			node.left = newNode(arr[i].(int))
			queue = append(queue, node.left)
		}
		i++

		// Process right child
		if i < len(arr) && arr[i] != nil {
			node.right = newNode(arr[i].(int))
			queue = append(queue, node.right)
		}
		i++
	}

	return root
}

func findPredecessor(root *tree, ele int) {
	var lastRight *tree
	var lastLeft *tree

	for root != nil {
		if root.val == ele {
			break
		}

		if ele < root.val {
			lastLeft = root
			root = root.left
		} else {
			lastRight = root
			root = root.right
		}
	}

	predRoot, succRoot := root, root
	pred := -1
	if predRoot.left == nil {
		if lastRight == nil {
			pred = -1
		} else {
			pred = lastRight.val
		}
	} else {
		predRoot = predRoot.left
		for predRoot != nil && predRoot.right != nil {
			predRoot = predRoot.right
		}
		pred = predRoot.val
	}

	succ := -1
	if succRoot.right == nil {
		if lastLeft == nil {
			succ = -1
		} else {
			succ = lastLeft.val
		}
	} else {
		succRoot = succRoot.right
		for succRoot != nil && succRoot.left != nil {
			succRoot = succRoot.left
		}
		succ = succRoot.val
	}

	fmt.Println("Predecessor of", ele, " is", pred)
	fmt.Println("Successor of", ele, " is", succ)
}

// Inorder Successor of a Node

// 	Case 1: Node has a right subtree
// 		Go to the right child, then keep going to the leftmost node from there.
// 	Case 2: No right subtree
// 		Go up the tree until you find a node that is the left child of its parent. That parent is the successor.

// Inorder Predecessor of a Node

// Case 1: Node has a left subtree
//
//	Go to the left child, then keep going to the rightmost node from there.
//
// Case 2: No left subtree
//
//	Go up the tree until you find a node that is the right child of its parent. That parent is the predecessor.
func main() {

	input := []interface{}{5, 2, 10, 1, 4, 7, 12}
	root := buildTreeFromArray(input)
	findPredecessor(root, 10)
	fmt.Println()
	findPredecessor(root, 12)

}
