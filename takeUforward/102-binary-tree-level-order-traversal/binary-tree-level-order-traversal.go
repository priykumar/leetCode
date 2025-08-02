/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {

    q:=[]*TreeNode{root}
    temp:=[]int{}
    res := [][]int{}
    if root == nil {
        return res
    }
    res=append(res, []int{root.Val})
    for len(q) > 0 {
        sz:=len(q)
        temp = []int{}
        for i:=0; i<sz; i++ {
            if q[i].Left != nil {
                q=append(q, q[i].Left)
                temp=append(temp, q[i].Left.Val)
            }
            if q[i].Right != nil {
                q=append(q, q[i].Right)
                temp=append(temp, q[i].Right.Val)
            }
        }
        if len(temp) > 0 {
            res=append(res, temp)
        }
        
        q=q[sz:]
    }
    
    return res
}