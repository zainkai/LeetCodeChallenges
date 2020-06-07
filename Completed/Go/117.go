/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type frame struct {
    root *Node
    depth int
}

func connect(root *Node) *Node {
    q := []frame{}
    if root != nil {
        q = append(q, frame{root, 0})
    }
    
    for len(q) > 0 {
        curr := q[0]
        q = q[1:]
        
        if len(q) > 0 && q[0].depth == curr.depth {
            curr.root.Next = q[0].root
        } else {
            curr.root.Next = nil
        }
        
        if curr.root.Left != nil {
            q = append(q, frame{ curr.root.Left, curr.depth+1})
        }
        if curr.root.Right != nil {
            q = append(q, frame{ curr.root.Right, curr.depth+1})
        }
    }
    
    return root
}
