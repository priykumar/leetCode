/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"

var res []int
func solve(root *TreeNode) {
    if root == nil {
        return 
    }
    solve(root.Left)
    res = append(res, root.Val)
    solve(root.Right)
}

func morris(root *TreeNode) {
    findPredecessor := func(node *TreeNode) *TreeNode {
        // No predecessor
        if node.Left == nil {
            return nil
        }

        node1 := node.Left
        for node1.Right != nil && node1.Right != node {
            node1 = node1.Right
        }

        return node1
    }

    curr := root
    for curr!=nil {
        if curr.Left == nil {
            res=append(res, curr.Val)
            curr=curr.Right
        } else {
            pre := findPredecessor(curr)
            // no link between curr and its predecessor, so establish a link
            if pre.Right == nil {
                pre.Right = curr
                curr = curr.Left
            } else {
                res=append(res, curr.Val)
                pre.Right=nil
                curr=curr.Right
            }
        }
    }
}

func inorderTraversal(root *TreeNode) []int {
    isMorris := true
    res = []int{}

    if !isMorris {
        solve(root)
    } else {
        fmt.Println("MORRIS TRAVERSAL")
        morris(root)
    }
    return res
}
