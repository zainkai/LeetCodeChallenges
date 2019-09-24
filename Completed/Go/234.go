/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	midNode := MidLL(head)

	a := ReverseLL(midNode)
	b := head

	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return true
}

func MidLL(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil && f.Next != nil {
		f = f.Next.Next
		s = s.Next
	}

	if f != nil {
		return s.Next
	}

	return s
}

func ReverseLL(head *ListNode) *ListNode {
	var prev, curr *ListNode = nil, head
	for curr != nil {
		t := curr.Next
		curr.Next = prev

		prev = curr
		curr = t
	}
	return prev
}