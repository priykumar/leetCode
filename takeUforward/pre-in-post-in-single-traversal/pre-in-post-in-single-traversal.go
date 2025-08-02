package main

import "fmt"

type tree struct {
	val   int
	left  *tree
	right *tree
}

var pre, in, post []int

func traversal(root *tree) {
	if root == nil {
		return
	}

	pre = append(pre, root.val)
	traversal(root.left)
	in = append(in, root.val)
	traversal(root.right)
	post = append(post, root.val)
}

func NewNode(value int) *tree {
	return &tree{
		val:   value,
		left:  nil,
		right: nil,
	}
}

// SetLeft sets the left child
func (node *tree) SetLeft(value int) *tree {
	node.left = NewNode(value)
	return node.left
}

// SetRight sets the right child
func (node *tree) SetRight(value int) *tree {
	node.right = NewNode(value)
	return node.right
}

func main() {
	root := NewNode(1)

	// Build tree structure:
	//       1
	//      / \
	//     2   3
	//    / \   \
	//   4   5   6

	root.SetLeft(2)
	root.SetRight(3)

	root.left.SetLeft(4)
	root.left.SetRight(5)

	root.right.SetRight(6)

	pre, in, post = []int{}, []int{}, []int{}
	traversal(root)
	fmt.Println("PREORDER:", pre)
	fmt.Println("INORDER:", in)
	fmt.Println("POSTORDER:", post)
}
