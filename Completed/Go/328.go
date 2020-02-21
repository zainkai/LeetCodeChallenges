/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil {
		return head
	}

	oddsHead := head
	evensHead := head.Next
	for curr, temp := head, head.Next; temp != nil; curr, temp = temp, temp.Next {
		curr.Next = temp.Next
	}

	oddsTail := oddsHead
	for ; oddsTail.Next != nil; oddsTail = oddsTail.Next {
	}
	oddsTail.Next = evensHead

	return oddsHead
}