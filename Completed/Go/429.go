/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 type frame struct {
    level int
    node *Node
}

func levelOrder(root *Node) [][]int {
    q := []frame{{1, root}}
    
    res := [][]int{}
    for len(q) > 0 {
        top := q[0]
        q = q[1:]
        
        if top.node == nil {
            continue
        } else if len(res) < top.level {
            res = append(res, []int{})
        }
        
        res[top.level-1] = append(res[top.level-1], top.node.Val)
        for _, n := range top.node.Children {
            q = append(q, frame{top.level+1, n})
        }
    }
    
    return res
}