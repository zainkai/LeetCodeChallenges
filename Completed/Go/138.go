/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    
	// step 1 weave list
	n := head
	for n != nil {
		n.Next = &Node{n.Val, n.Next, nil}
		n = n.Next.Next
	}
    
    

	// step 2 assign random points
	n = head
	for n != nil {
        if n.Random != nil {
           n.Next.Random = n.Random.Next 
        }
		n = n.Next.Next
	}

	// step 3 unweave lists
    HeadPrime := head.Next
    
	n = head
	nPrime := HeadPrime
	for n != nil {
		n.Next = n.Next.Next
        if nPrime.Next != nil {
            nPrime.Next = nPrime.Next.Next
        }
        
        n = n.Next
        nPrime = nPrime.Next
	}

	return HeadPrime
}
