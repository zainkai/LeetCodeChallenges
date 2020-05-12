/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }
    
    t := 0
    for _, child := range root.Children {
        res := maxDepth(child)
        if res > t {
            t = res
        }
    }
    
    return t+1
}