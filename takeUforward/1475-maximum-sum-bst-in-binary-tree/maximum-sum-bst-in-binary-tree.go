/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBST(root *TreeNode, res *int) (int, int, int, bool){
    if root == nil {
        return int(-1e9), int(1e9), 0, true // max, min, sum, isbst
    }

    lmax, lmin, lsum, lbst := isBST(root.Left, res)
    rmax, rmin, rsum, rbst := isBST(root.Right, res)

    isbst := true
    if !(lmax < root.Val && root.Val < rmin) {
        isbst = false
    }

    isbst = isbst && lbst && rbst
    maxval := max(max(lmax, rmax), root.Val)
    minval := min(min(lmin, rmin), root.Val)
    sum := root.Val + lsum + rsum

    if isbst {
        *res = max(*res, sum)
    }

    // fmt.Println("For", root.Val, "=", maxval, minval, sum, isbst)
    return maxval, minval, sum, isbst
}
func maxSumBST(root *TreeNode) int {
    res := 0
    isBST(root, &res)
    return res
}