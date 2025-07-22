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
    if root == nil {
        return 
    }
    solve(root.Left)
    res = append(res, root.Val)
    solve(root.Right)
}
func inorderTraversal(root *TreeNode) []int {
    res = []int{}
    solve(root)
    return res
}