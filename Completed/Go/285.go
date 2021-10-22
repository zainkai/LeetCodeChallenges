/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
    found := false
    return helper(root, p, &found)
}

func helper(root, p *TreeNode, found *bool) *TreeNode {
    if root == nil {
        return nil
    }
    
    // left
    if n := helper(root.Left, p, found); n != nil {
        return n
    }
    
    // center
    if root == p {
        *found = true
    } else if *found {
        return root
    }
    
    // right
    return helper(root.Right, p, found)
}
