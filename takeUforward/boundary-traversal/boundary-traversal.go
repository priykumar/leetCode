package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

var leafNode []*tree

// isLeaf checks if a node is a leaf
func isLeaf(node *tree) bool {
	return node != nil && node.left == nil && node.right == nil
}

// getLeftBoundary gets left boundary nodes (excluding leaves)
func getLeftBoundary(root *tree, result *[]int) {
	if root == nil || isLeaf(root) {
		return
	}

	*result = append(*result, root.val)

	if root.left != nil {
		getLeftBoundary(root.left, result)
	} else {
		getLeftBoundary(root.right, result)
	}
}

// getLeafNodes gets all leaf nodes from left to right
func getLeafNodes(root *tree, result *[]int) {
	if root == nil {
		return
	}

	if isLeaf(root) {
		*result = append(*result, root.val)
		return
	}

	getLeafNodes(root.left, result)
	getLeafNodes(root.right, result)
}

// getRightBoundary gets right boundary nodes (excluding leaves) in reverse order
func getRightBoundary(root *tree, result *[]int) {
	if root == nil || isLeaf(root) {
		return
	}

	if root.right != nil {
		getRightBoundary(root.right, result)
	} else {
		getRightBoundary(root.left, result)
	}

	*result = append(*result, root.val)
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

func main() {
	// Input: [1, 2, null, 4, 9, 6, 5, 3, null, null, null, null, null, 7, 8]
	input := []interface{}{1, 2, nil, 4, 9, 6, 5, 3, nil, nil, nil, nil, nil, 7, 8}

	root := buildTreeFromArray(input)
	result := []int{}
	if !isLeaf(root) {
		result = append(result, root.val)
	}
	getLeftBoundary(root.left, &result)
	getLeafNodes(root, &result)
	getRightBoundary(root.right, &result)

	fmt.Println(result)
}
