/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findNextGreaterElement(a []int, s, e, ele int) int {
    fmt.Println(a[s:e+1], ele)
    for i:=s; i<=e; i++ {
        if a[i]>ele {
            return i
        }
    }
    
    return e+1
}

func solve(preorder []int, start, end int) *TreeNode {
    if start>end {
        return nil
    }

    root := &TreeNode{preorder[start], nil, nil}
    if start == end {
        return root
    }
    idx := findNextGreaterElement(preorder, start, end, preorder[start])
    root.Left = solve(preorder, start+1, idx-1)
    root.Right = solve(preorder, idx, end)

    return root
}

func bstFromPreorder(preorder []int) *TreeNode {
    return solve(preorder, 0, len(preorder)-1)
}