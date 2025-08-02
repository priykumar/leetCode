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
        return 
    }

    if root == p {
        *lca = p
        return 
    }
    if root == q {
        *lca =q
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
    solve(root, p, q, &lca)
    return lca
}