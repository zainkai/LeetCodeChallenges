/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a, b := headA, headB
    for {
        if a == b {
            return a
        }
        
        if a == nil {
            a = headB
        } else {
            a = a.Next
        }
        
        if b == nil {
            b = headA
        } else {
            b = b.Next
        }
    }
}