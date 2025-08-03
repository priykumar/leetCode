/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findPrdecessor(root *TreeNode) *TreeNode {
    if root.Left == nil {
        return nil
    }

    x:=root.Left
    for x.Right != nil && x.Right!=root{
        x = x.Right
    }

    return x
}

func morrisInorder(root *TreeNode, k int) int {
    curr := root
    count:=0
    for curr != nil {
        if curr.Left == nil {
            count++
            // fmt.Println(count, curr)
            if count == k {
                return curr.Val
            }
            curr=curr.Right
        } else {
            pre := findPrdecessor(curr)
            if pre.Right == nil {
                pre.Right = curr
                // fmt.Println(pre, pre.Right)
                curr = curr.Left
            } else {
                count++
                //fmt.Println(count, curr)
                if count == k {
                    return curr.Val
                }
                pre.Right = nil
                curr = curr.Right
            }
        }
    }
    return -1
}

func kthSmallest(root *TreeNode, k int) int {
    return morrisInorder(root, k)
}