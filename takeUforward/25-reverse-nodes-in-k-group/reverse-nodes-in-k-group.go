/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy := &ListNode{0, head}
    var start_prev, start, end, end_next, curr *ListNode
    start, curr, start_prev = head, head, dummy
    var prev, next *ListNode

    count := 1
    for curr!=nil{
        if count % k == 0 {
            // reverse
            start_prev = end
            if start_prev == nil {
                start_prev = dummy
            }
            start = end_next
            if start == nil {
                start = head
            }

            end = curr
            end_next = curr.Next
            curr=curr.Next
            count++
            
            // fmt.Println(count, start_prev, start, end, end_next, curr)
            c_curr := start
            for c_curr != end_next {
                next = c_curr.Next
                c_curr.Next = prev
                prev=c_curr
                c_curr=next
            }

            // fmt.Println(count, start_prev, start, end, end_next)
            start_prev.Next = end
            start.Next = end_next
            // fmt.Println(count, start_prev, start, end, end_next)
            // fmt.Println()
            end=start
        } else {
            curr = curr.Next
            count++
        }
    }

    return dummy.Next
}