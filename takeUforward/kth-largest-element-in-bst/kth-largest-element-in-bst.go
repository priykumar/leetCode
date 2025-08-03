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

func findPredecessor(root *tree) *tree {
	x := root.left
	for x.right != nil && x.right != root {
		x = x.right
	}

	return x
}

func kth_largest(root *tree, k int) int {
	// MORRIS INORDER
	curr := root
	res := []int{}
	for curr != nil {
		if curr.left == nil {
			res = append(res, curr.val)
			curr = curr.right
		} else {
			pre := findPredecessor(curr)
			if pre.right == nil {
				pre.right = curr
				curr = curr.left
			} else {
				res = append(res, curr.val)
				pre.right = nil
				curr = curr.right
			}
		}
	}

	sz := len(res)
	fmt.Println(res)
	return res[sz-k]
}

func main() {
	input := []interface{}{3, 1, 4, nil, 2}
	root := buildTreeFromArray(input)
	fmt.Println(kth_largest(root, 1))

	fmt.Println()

	input = []interface{}{5, 3, 6, 2, nil, nil, nil, 1}
	root = buildTreeFromArray(input)
	fmt.Println(kth_largest(root, 3))

}
