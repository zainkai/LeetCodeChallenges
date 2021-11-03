/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func cloneTree(root *Node) *Node {
    if root == nil {
        return nil
    }
    
    cRoot := cloneNode(root)
    
    for i, child := range root.Children {
        cRoot.Children[i] = cloneTree(child)
    }
    
    return cRoot
}

func cloneNode(root *Node) *Node {    
    return &Node{
        root.Val,
        make([]*Node, len(root.Children)),
    }
}
