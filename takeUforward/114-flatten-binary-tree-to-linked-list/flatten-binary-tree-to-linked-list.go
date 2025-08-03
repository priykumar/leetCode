/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
    curr := root
    for curr != nil {
        temp:=curr.Right
        curr.Right = curr.Left
        curr.Left = nil
        fmt.Println(curr)
        
        x:=curr
        for x.Right != nil {
            x = x.Right
        }
        x.Right = temp

        curr = curr.Right
    }
}