/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode, balanced *bool) int {
    if root==nil {
        return -1
    }

    l:=height(root.Left, balanced)
    r:=height(root.Right, balanced)

    if l-r>1 || r-l>1 {
        *balanced = (*balanced) && false
    }

    return max(l, r)+1
}
func isBalanced(root *TreeNode) bool {
    balanced := true
    height(root, &balanced)
    return balanced
}