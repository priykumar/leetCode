/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var lca *TreeNode
func solve(root, p, q *TreeNode) (bool, bool){
    if root == nil || lca != nil {
        return false, false
    }

    lpFound, lqFound := solve(root.Left, p, q)
    rpFound, rqFound := solve(root.Right, p, q)

    if ((lpFound && rqFound) || (lqFound && rpFound)) && lca == nil {
        lca = root
    }

    pFound, qFound := lpFound||rpFound, lqFound||rqFound
    if root == p {
        pFound = true
        if qFound && lca == nil {
            lca = p
        }
    }
    if root == q {
        qFound = true
        if pFound && lca == nil {
            lca = q
        }
    }

    return pFound, qFound
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
  lca = nil
  solve(root, p, q)
  fmt.Println(lca)
  return lca
}