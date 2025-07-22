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
    solve(root.Left)
    solve(root.Right)
    res=append(res, root.Val)
}
func postorderTraversal(root *TreeNode) []int {
    res = []int{}
    solve(root)
    return res
}