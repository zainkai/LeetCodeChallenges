/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	t := head
	for t != nil && t.Next != nil {
		if t.Val == t.Next.Val {
			t.Next = t.Next.Next
		} else {
			t = t.Next
		}
	}
	return head
}