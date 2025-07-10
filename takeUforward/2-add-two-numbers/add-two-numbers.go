/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    carry:=0
    dummy:=&ListNode{0, nil}
    curr, sum := dummy, 0
    for l1 != nil && l2 != nil {
        sum = l1.Val + l2.Val + carry
        if sum >= 10 {
            carry=1
            sum=sum%10
        } else {
            carry=0
        }
        curr.Next=&ListNode{sum, nil}
        curr=curr.Next
        l1=l1.Next
        l2=l2.Next
    }

    if l1 != nil {
        for l1 != nil {
            sum = l1.Val + carry
            if sum >= 10 {
                carry=1
                sum=sum%10
            } else {
                carry=0
            }
            curr.Next=&ListNode{sum, nil}
            curr=curr.Next
            l1=l1.Next
        }
    } else if l2 != nil {
        for l2 != nil {
            sum = l2.Val + carry
            if sum >= 10 {
                carry=1
                sum=sum%10
            } else {
                carry=0
            }
            curr.Next=&ListNode{sum, nil}
            curr=curr.Next
            l2=l2.Next
        }
    }

    if carry == 1 {
        curr.Next=&ListNode{1, nil}
    }

    return dummy.Next
}