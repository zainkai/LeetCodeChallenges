/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    // Divide List in half
    left, right := halfLL(head)
    
    left = mergeSort(left)
    right = mergeSort(right)
    
    
    tmpHead := &ListNode{}
    curr := tmpHead
    for left != nil && right != nil {
        if left.Val < right.Val {
            curr.Next = left
            curr = curr.Next
            left = left.Next
        } else {
            curr.Next = right
            curr = curr.Next
            right = right.Next
        }
    }
    if left != nil {
        curr.Next = left
    }
    if right != nil {
        curr.Next = right
    }
    
    return tmpHead.Next
}

func halfLL(head *ListNode) (*ListNode, *ListNode) {
    slow, fast := head, head
    prev := slow
    for fast != nil && fast.Next != nil {
        prev = slow
        slow = slow.Next
        
        fast = fast.Next.Next
    }
    prev.Next = nil
    
    return head, slow
}
