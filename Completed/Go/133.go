/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

 func cloneGraph(node *Node) *Node {
    nodeMap := make(map[*Node]*Node)
    
    return helper(node, nodeMap)
}

func helper(oldNode *Node, nodeMap map[*Node]*Node) *Node {
    n, visited := cloneOrGetNode(oldNode, nodeMap)
    if visited {
        return n
    }
    
    for _, conn := range oldNode.Neighbors {
        n.Neighbors = append(n.Neighbors, helper(conn, nodeMap))
    }
    
    return n
}

func cloneOrGetNode(n *Node, nodeMap map[*Node]*Node) (*Node, bool) {
    if n == nil {
        return nil, true
    }
    
    _, ok := nodeMap[n]
    if !ok {
        nodeMap[n] = &Node{
            n.Val,
            []*Node{},
        }
    }
    
    return nodeMap[n], ok
}