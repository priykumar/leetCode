/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    var end *ListNode
    curr, len := head, 0
    for curr != nil {
        end=curr
        curr=curr.Next
        len++
    }

    if len == 0{
        return nil
    }

    k=k%len
    if k == 0 {
        return head
    }

    breakLinkAfter := len-k
    i:=1
    curr=head
    for i<breakLinkAfter {
        curr = curr.Next
        i++
    }

    start := curr.Next
    curr.Next=nil
    end.Next=head   

    return start
}