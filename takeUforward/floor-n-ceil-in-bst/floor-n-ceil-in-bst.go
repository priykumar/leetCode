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

func floor_n_ceil(root *tree, ele int) {
	var lastRight *tree = &tree{-1, nil, nil}
	var lastLeft *tree = &tree{-1, nil, nil}

	found := false

	for root != nil {
		if root.val == ele {
			found = true
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

	if found {
		fmt.Println("Floor and Ceil of", ele, "is: ", ele, ele)
	} else {
		fmt.Println("Floor and Ceil of", ele, "is: ", lastRight.val, lastLeft.val)
	}
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

	input := []interface{}{8, 4, 12, 2, 6, 10, 14}
	root := buildTreeFromArray(input)
	floor_n_ceil(root, 10)
	floor_n_ceil(root, 12)
	floor_n_ceil(root, 11)
	floor_n_ceil(root, 15)

}
