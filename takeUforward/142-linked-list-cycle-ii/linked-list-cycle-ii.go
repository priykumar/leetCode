/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    slow, fast := head, head.Next

    for fast!=nil && fast.Next != nil && slow != fast {
        slow = slow.Next
        fast = fast.Next.Next
        fmt.Println(slow, fast)
    }

    
    if fast == nil || fast.Next == nil {
        return nil
    }

    fast = head
    fmt.Println()
    for slow != fast && slow.Next != fast {
        slow=slow.Next
        fast=fast.Next
    }

    return fast
}