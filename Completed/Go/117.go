/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    
    q := []*Node{root}
    for len(q) > 0 {
        stop := len(q)
        for i := 0; i < stop; i++ {
            top := q[0]
            q = q[1:]
            
            if i < stop-1 {
                top.Next = q[0]
            }
            if top.Left != nil {
                q = append(q, top.Left)
            }
            if top.Right != nil {
                q = append(q, top.Right)
            }
        }
    }
    
    return root
}
