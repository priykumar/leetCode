/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(lTree, rTree *TreeNode) bool {
    if lTree == nil && rTree == nil {
        return true
    }
    if (lTree == nil && rTree != nil) || (lTree != nil && rTree == nil) {
        // fmt.Println("here1")
        return false
    }

    l := solve(lTree.Left, rTree.Right)
    r := solve(lTree.Right, rTree.Left)

    // fmt.Println(l,r,lTree.Val == rTree.Val, lTree.Val, rTree.Val)
    return l && r && lTree.Val == rTree.Val

}
func isSymmetric(root *TreeNode) bool {
    return solve(root.Left, root.Right)
}