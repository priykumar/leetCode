/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func hasCycle(head *ListNode) bool {
//     m:=make(map[*ListNode]bool)
//     for head!=nil {
//         if _, exist := m[head]; exist {
//             return true
//         }
//         m[head]=true
//         head=head.Next
//     }
//     return false
// }

func hasCycle(head *ListNode) bool {
    for head!=nil {
        if head.Val == 1<<32 {
            return true
        }
        head.Val = 1<<32
        head=head.Next
    }
    return false
}