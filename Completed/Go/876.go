/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	count := 0
	t := head
	for t != nil {
		t = t.Next
		count += 1
	}

	for i := 0; i < (count / 2); i++ {
		head = head.Next
	}

	return head
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	s := head
	f := head
	for f != nil && f.Next != nil {
		s = s.Next
		f = f.Next.Next
	}
	return s
}
