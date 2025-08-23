/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    curr := head
    length := 0
    for curr!=nil {
        length++
        curr=curr.Next
    }

    nodeFromFront := length-n
    if nodeFromFront == 0 {
        return head.Next
    }

    curr = head
    count:=0
    for curr != nil{
        count++
        if count == nodeFromFront {
            curr.Next=curr.Next.Next
            break
        }
        curr=curr.Next
    }

    return head
}