/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(root *TreeNode, isBST *bool) (int, int) {
    if root == nil {
        return math.MaxInt64, math.MinInt64 // int(1e9), int(-1e9)
    }

    lmin, lmax := solve(root.Left, isBST)
    rmin, rmax := solve(root.Right, isBST)

    if !(lmax < root.Val && root.Val < rmin) {
        *isBST = *isBST && false
    }

    fmt.Println(root.Val, min(root.Val, min(lmin, rmin)), max(root.Val, min(lmax, rmax)))
    return min(root.Val, min(lmin, rmin)), max(root.Val, max(lmax, rmax))
}
func isValidBST1(root *TreeNode) bool {
    isBST := true
    solve(root, &isBST)
    return isBST
}

func inOrder(root *TreeNode, a *[]int) {
    if root==nil {
        return
    }
    inOrder(root.Left, a)
    *a=append(*a, root.Val)
    inOrder(root.Right, a)
}
func isValidBST(root *TreeNode) bool {
    a:=[]int{}
    inOrder(root, &a)
    
    if len(a)<=1{
        return true
    }
    for i:=1; i<len(a); i++ {
        if a[i-1]>=a[i] {
            return false
        }
    }

    return true
}