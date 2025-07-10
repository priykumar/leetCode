/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    // slow, fast := head, head

    // for fast != nil && fast.Next != nil {
    //     slow=slow.Next
    //     fast=fast.Next.Next

    //     if slow == fast {
    //         return true
    //     }
    // }
    if head == nil {
        return false
    }

    slow, fast := head, head.Next

    for slow != fast {
        for fast == nil || fast.Next == nil {
            return false
        }
        slow = slow.Next
        fast = fast.Next.Next
    }

    return true
}