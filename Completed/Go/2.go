/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var front *ListNode = &ListNode{0, nil}

	carry := 0
	curr := front
	for l1 != nil || l2 != nil {
		a, b := 0, 0
		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry
		carry = 0
		if sum > 9 {
			sum = sum % 10
			carry = 1
		}

		curr.Next = &ListNode{sum, nil}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}

	return front.Next
}