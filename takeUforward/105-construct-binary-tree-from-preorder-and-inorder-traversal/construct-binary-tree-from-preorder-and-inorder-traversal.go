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

func solve(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode{
    if inStart > inEnd {
        return nil
    }

    root := &TreeNode{preorder[preStart], nil, nil}
    idx := findInInorder(inorder, inStart, inEnd, preorder[preStart])

    count1 := idx-inStart
    root.Left = solve(preorder, inorder, preStart+1, preStart+count1, inStart, idx-1)
    count2 := inEnd-idx
    root.Right = solve(preorder, inorder, preEnd-count2+1, preEnd, idx+1, inEnd)

    return root
}
func buildTree(preorder []int, inorder []int) *TreeNode {
    return solve(preorder, inorder, 0, len(inorder)-1, 0, len(inorder)-1,)
}