/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head.Next == nil {
        return true
    }

    var prev, l1, l2 *ListNode
    slow, fast := head, head
    for fast.Next != nil && fast.Next.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    l1 = slow
    l2 = slow.Next
    isOdd:=false
    if fast.Next == nil {
        isOdd = true
    }
    slow.Next = nil

    var curr, next *ListNode
    curr, prev = head, nil
    for curr != nil {
        next = curr.Next
        curr.Next=prev
        prev=curr
        curr=next
    }

    if isOdd {
        l1 = l1.Next
    }

    for l1 != nil && l2 != nil {
        if l1.Val != l2.Val {
            return false
        }
        l1=l1.Next
        l2=l2.Next
    }

    if l1==nil && l2==nil {
        return true
    }
    return false
}