/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func height(root *TreeNode) int {
    if root == nil {
        return -1
    }

    l:=height(root.Left)
    r:=height(root.Right)

    return max(l, r) + 1
}

func maxDepth(root *TreeNode) int {
    return height(root)+1 // +1 is added becuase we are calculating depth, for height calculation there should't be +1
}