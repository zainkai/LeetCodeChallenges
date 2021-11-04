/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func findRoot(tree []*Node) *Node {
    refCounter := map[*Node]int{}
    for _, n := range tree {
        if n == nil {
            continue
        }
        
        for _, child := range n.Children {
            refCounter[child]++
        }
    }
    
    for _, n := range tree {
        if refCounter[n] == 0 {
            return n
        }
    }
    
    return nil
}
