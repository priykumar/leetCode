/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(root, p, q *TreeNode, lca **TreeNode) {
    if root == nil || *lca != nil{
        return 
    }

    if p.Val < root.Val && root.Val < q.Val {
        *lca = root
        fmt.Println("here1")
        return 
    }

    if root == p {
        *lca = p
        fmt.Println("here2")
        return 
    }
    if root == q {
        *lca =q
        fmt.Println("here3")
        return 
    }

    if p.Val < root.Val && q.Val < root.Val {
        solve(root.Left, p, q, lca)
    } else {
        solve(root.Right, p, q, lca)
    }

}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
        p, q = q, p
    }

    var lca *TreeNode
    fmt.Println(lca)
    solve(root, p, q, &lca)
    fmt.Println(lca)
    return lca
}