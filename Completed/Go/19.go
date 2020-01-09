func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var sentinel *ListNode = new(ListNode)
	sentinel.Next = head

	var pacer, follower *ListNode = head, sentinel
	for i := 0; i < n; i++ {
		pacer = pacer.Next
	}
	for pacer != nil {
		pacer = pacer.Next
		follower = follower.Next
	}

	follower.Next = follower.Next.Next

	return sentinel.Next
}

/** OLD algo

func removeNthFromEnd(head *ListNode, n int) *ListNode {
  var sentinel *ListNode = &ListNode{ 0, head }
  
  var pacer, follower *ListNode = head, nil
  n--
  for pacer != nil {
    if n == 0 {
      follower = sentinel
    } else if n < 0 {
      follower = follower.Next
    }
    
    pacer = pacer.Next
    n--
  }
  
  follower.Next = follower.Next.Next
  
  
  return sentinel.Next
}
*/