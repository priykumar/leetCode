/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//     sz := 0
//     curr := head
//     for curr != nil {
//         sz++
//         curr=curr.Next
//     }

//     posFromStart := sz-n+1
//     if posFromStart == 1 {
//         return head.Next
//     }

//     currPos:=1
//     curr=head
//     var prev *ListNode

//     for currPos < posFromStart {
//         currPos++
//         prev=curr
//         curr=curr.Next
//     }

//     prev.Next = curr.Next
//     return head
// }

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    dummy := &ListNode{0, head}
    slow, fast := dummy, dummy

    for i:=0; i<=n; i++ {
        fast=fast.Next
    }

    for fast!=nil {
        slow=slow.Next
        fast=fast.Next
    }

    slow.Next = slow.Next.Next
    return dummy.Next
}