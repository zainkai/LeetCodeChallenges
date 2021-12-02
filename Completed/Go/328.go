/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    
    var oddHead, oddTail *ListNode = nil, nil
    var evenHead, evenTail *ListNode = nil, nil
    isOdd := true
    
    oddHead = head
    oddTail = oddHead
    head = head.Next
    if head != nil {
        evenHead = head
        evenTail = evenHead
        head = head.Next
    }
    
    for head != nil {
        if isOdd {
            oddTail.Next = head
            oddTail = oddTail.Next
        } else {
            evenTail.Next = head
            evenTail = evenTail.Next
        }
        isOdd = !isOdd
        head = head.Next
    }
    
    oddTail.Next = evenHead
    evenTail.Next = nil
    
    return oddHead
}
