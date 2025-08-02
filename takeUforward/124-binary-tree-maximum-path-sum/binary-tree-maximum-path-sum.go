/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var dp map[*TreeNode]int
func solve(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    
    if v, exist := dp[root]; exist {
        return v
    }

    lsum := solve(root.Left, res)
    rsum := solve(root.Right, res)

    if *res < root.Val + lsum + rsum {
        *res = root.Val + lsum + rsum
    }

    tempRes := max(max(lsum, rsum), 0)+root.Val
    *res = max(*res, tempRes)

    // *res = max(*res, root.Val + lsum + rsum)
    fmt.Println(root.Val, *res, lsum, rsum)
    // dp[root] = *res
    return tempRes
}
func maxPathSum(root *TreeNode) int {
    dp = make(map[*TreeNode]int)
    res := int(-1e9)
    solve(root, &res)

    return res
}