/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func plusOne(head *ListNode) *ListNode {
    carry := helper(head)
    if carry == 1 {
        return &ListNode{
            Val: 1,
            Next: head,
        }
    }
    return head
}

func helper (n *ListNode) int {
    if n == nil {
        return 1
    }
    
    carry := helper(n.Next)
    n.Val += carry
    if n.Val >= 10 {
        n.Val = n.Val % 10
        return 1
    }
    return 0
}
