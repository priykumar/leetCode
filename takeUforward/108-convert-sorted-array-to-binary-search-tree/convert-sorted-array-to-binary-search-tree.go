/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func solve(nums []int, start, end int) *TreeNode{
    fmt.Println(start, end)
    if start > end {
        return nil
    }
    
    mid := (end+start)/2
    root := &TreeNode{nums[mid], nil, nil}
    root.Left = solve(nums, start, mid-1)
    root.Right = solve(nums, mid+1, end)

    return root

}
func sortedArrayToBST(nums []int) *TreeNode {
    return solve(nums, 0, len(nums)-1)
}