/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow := head
	fast := head

	for slow != nil && fast != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}

		if slow != nil && fast != nil && slow.Val == fast.Val {
			return true
		}
	}
	return false
}