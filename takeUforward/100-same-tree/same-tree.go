/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }
    if (p == nil && q != nil) || (p != nil && q == nil) || (p.Val != q.Val) {
        return false
    }

    l := solve(p.Left, q.Left)
    r := solve(p.Right, q.Right)

    return l && r
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return solve(p, q)
}