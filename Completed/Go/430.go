func flatten(root *Node) *Node {
	first, _ := dfs(root)
	return first
}

func dfs(root *Node) (*Node, *Node) {    
	tmp, last := root, root
    
	for tmp != nil {
		if tmp.Child != nil {
			first, last := dfs(tmp.Child)
            
            tmp.Child = nil // requires child is nil'ed
            
            tmpNext := tmp.Next
			connect(tmp, first)
			connect(last, tmpNext)
			tmp = last
		}
        
        last = tmp
		tmp = tmp.Next
	}

	return root, last
}

func connect(first, second *Node) {
    if first != nil {
       first.Next = second 
    }
    
    if second != nil {
       second.Prev = first 
    }
}