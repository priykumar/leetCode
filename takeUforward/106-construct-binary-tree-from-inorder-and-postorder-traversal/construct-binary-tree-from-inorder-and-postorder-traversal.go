/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findInInorder(inorder []int, inStart, inEnd, ele int) int {
    for i:=inStart; i<=inEnd; i++ {
        if inorder[i] == ele {
            return i
        }
    }

    return -1
}
func solve(inorder, postorder []int, inStart, inEnd, postStart, postEnd int) *TreeNode {
    if inStart > inEnd {
        return nil
    }

    fmt.Println(inStart, inEnd, postStart, postEnd)
    
    root := &TreeNode{postorder[postEnd], nil, nil}
    idx := findInInorder(inorder, inStart, inEnd, postorder[postEnd])

    count1:=idx-inStart
    root.Left = solve(inorder, postorder, inStart, idx-1, postStart, postStart+count1-1)
    count2:=inEnd-(idx)
    root.Right = solve(inorder, postorder, idx+1, inEnd, postEnd-count2, postEnd-1)

    return root
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    return solve(inorder, postorder, 0, len(inorder)-1, 0, len(inorder)-1)
}