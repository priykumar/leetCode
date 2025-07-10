/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var curr, prev, next *ListNode
    curr = head

    for curr!=nil {
        next=curr.Next
        curr.Next=prev
        prev=curr
        curr=next
    }

    return prev
}