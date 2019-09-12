/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	temp := &ListNode{}
	result := temp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			temp.Next = l1
			l1 = l1.Next
		} else {
			temp.Next = l2
			l2 = l2.Next
		}

		temp = temp.Next
	}

	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp.Next = l2
			l2 = l2.Next
		} else {
			temp.Next = l1
			l1 = l1.Next
		}

		temp = temp.Next
	}

	return result.Next
}