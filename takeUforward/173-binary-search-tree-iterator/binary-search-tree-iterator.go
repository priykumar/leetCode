/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
    inorder []int
    currPos int
}

func (b *BSTIterator) inOrder(root *TreeNode) {
    if root == nil {
        return 
    }
    b.inOrder(root.Left)
    b.inorder = append(b.inorder, root.Val)
    b.inOrder(root.Right)
}

func Constructor(root *TreeNode) BSTIterator {
    bst := BSTIterator{
        inorder: []int{},
        currPos: 0,
    }
    bst.inOrder(root)

    return bst
}


func (this *BSTIterator) Next() int {
    val := this.inorder[this.currPos]
    this.currPos++
    return val
}


func (this *BSTIterator) HasNext() bool {
    if this.currPos<len(this.inorder) {
        return true
    }
    return false
}


/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */