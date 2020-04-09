/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
    trav := []int{}
    helper(root, &trav)
    
    return trav
}

func helper(root *Node, trav *[]int) {
    if root == nil {
        return
    }
    
    *trav = append(*trav, root.Val)
    for _, child := range root.Children {
        helper(child, trav)
    }
}