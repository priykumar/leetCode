/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []int
func solve(root *TreeNode) {
    if root==nil {
        return 
    }
    res=append(res, root.Val)
    solve(root.Left)
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
    for curr != nil {
        if curr.Left == nil {
            res = append(res, curr.Val)
            curr = curr.Right
        } else {
            pre := findPredecessor(curr)
            if pre.Right == nil {
                res = append(res, curr.Val)
                pre.Right = curr
                curr = curr.Left
            } else {
                pre.Right = nil
                curr = curr.Right
            }
        }
    }
}
func preorderTraversal(root *TreeNode) []int {
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