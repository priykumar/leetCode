/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
    curr := root
    for curr != nil {
        if curr.Val == val {
            return curr
        } 
        if val < curr.Val {
            curr = curr.Left
        } else if val > curr.Val {
            curr = curr.Right
        }
    }

    return nil
}