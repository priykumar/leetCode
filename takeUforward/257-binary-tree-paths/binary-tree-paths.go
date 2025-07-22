/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res []string
func solve(curr *TreeNode, temp string) {
    if curr.Left == nil && curr.Right == nil {
        temp = fmt.Sprintf("%s%d",temp, curr.Val)
        res = append(res, temp)
        return 
    }

    prevTemp := temp
    temp = fmt.Sprintf("%s%d->",temp, curr.Val)
    if curr.Left != nil {
        solve(curr.Left, temp)
    }
    if curr.Right != nil {
        solve(curr.Right, temp)
    }
    temp=prevTemp
}
func binaryTreePaths(root *TreeNode) []string {
    //temp:=fmt.Sprintf("%d", root.Val)
    res = []string{}
    solve(root, "")
    return res
}