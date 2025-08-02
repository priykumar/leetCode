/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    q := []*TreeNode{root}
    res := [][]int{}
    // temp := []int{}
    if root == nil {
        return res
    }
    level := 1
    for len(q) > 0 {
        sz := len(q)
        level++
        temp := []int{}
        for i:=0; i<sz; i++ {
            temp = append(temp, q[i].Val)
            if q[i].Left != nil {
                q = append(q, q[i].Left)
            }
            if q[i].Right != nil {
                q = append(q, q[i].Right)
            }
        }
        q=q[sz:]

        if level % 2 == 1 {
            l, r := 0, len(temp)-1
            for l<r {
                temp[l], temp[r] = temp[r], temp[l]
                l++
                r--
            }
        }
        
        res = append(res, temp)
    }

    return res
}