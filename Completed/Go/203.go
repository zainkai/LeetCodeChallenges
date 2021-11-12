/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    sentinel := &ListNode{
        0,
        head,
    }
    
    tmp := sentinel
    for tmp != nil && tmp.Next != nil {
        if tmp.Next.Val == val {
            tmp.Next = tmp.Next.Next // same as delete
        } else {
            tmp = tmp.Next
        }   
    }
    
    return sentinel.Next
}
