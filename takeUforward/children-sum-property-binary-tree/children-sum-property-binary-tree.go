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

func solve(root *tree, res *bool) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return root.val
	}

	lsum := solve(root.left, res)
	rsum := solve(root.right, res)

	if root.val != lsum+rsum {
		*res = false
	}

	return root.val
}

func main() {
	root := buildTreeFromArray([]interface{}{1, 4, 3, 5})
	res := true
	solve(root, &res)
	fmt.Println(res)

	root = buildTreeFromArray([]interface{}{10, 4, 6, 1, 3, 2, 4})
	res = true
	solve(root, &res)
	fmt.Println(res)
}
