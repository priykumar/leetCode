/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    // SC : O(n)
    // m := map[*ListNode]bool{}

    // for headA != nil {
    //     m[headA] = true
    //     headA=headA.Next
    // }

    // for headB != nil {
    //     if m[headB] {
    //         return headB
    //     }
    //     headB=headB.Next
    // }

    // return nil


    lenA, lenB := 0, 0
    currA, currB := headA, headB

    for currA != nil {
        lenA++
        currA=currA.Next
    }

    for currB != nil {
        lenB++
        currB=currB.Next
    }

    currA, currB = headA, headB
    if lenA > lenB {
        i := 0
        for i < lenA-lenB {
            i++
            currA=currA.Next
        }
    } else if lenB > lenA {
        i := 0
        for i < lenB-lenA {
            i++
            currB=currB.Next
        }
    }

    for currA != currB {
        currA = currA.Next
        currB = currB.Next
    }

    return currA
}