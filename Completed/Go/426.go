/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 * }
 */

// double pointers are hard to deal with so this is easier
type data struct {
    pred *Node
    head *Node
}

func treeToDoublyList(root *Node) *Node {
    if root == nil {
        return nil
    }
    
    
    d := &data{nil, nil}
    helper(root, d)
    
    d.head.Left = d.pred
    d.pred.Right = d.head
    
    return d.head
}

func helper(root *Node, d *data) {
    // traverse left
    if root.Left != nil {
        helper(root.Left, d)
    }
    
    // curr
    if d.pred != nil {
        d.pred.Right = root
        root.Left = d.pred
    } else {
        root.Left = nil
        d.head = root
    }
    d.pred = root
    
    //traverse right
    if root.Right != nil {
        helper(root.Right, d)
    }
}
