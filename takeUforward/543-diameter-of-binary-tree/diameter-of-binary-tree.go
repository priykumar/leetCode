/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode, diameter *int) int {
    if root == nil {
        return -1
    }

    l := height(root.Left, diameter)
    r := height(root.Right, diameter)
    fmt.Println(root.Val, l, r)

    *diameter = max(*diameter, l + r+1)
    return max(l,r) + 1
}
func diameterOfBinaryTree(root *TreeNode) int {
    diameter := -1
    if root.Left == nil && root.Right == nil {
        return 0
    }
    height(root, &diameter)
    return diameter+1
}