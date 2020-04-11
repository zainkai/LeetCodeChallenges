/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pruneTree(root *TreeNode) *TreeNode {
    if sum(root) == 0 {
        return nil
    }
    return root
}

func sum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    
    left := sum(root.Left)
    if left == 0 {
        root.Left = nil
    }
    
    right := sum(root.Right)
    if right == 0 {
        root.Right = nil
    }
    
    return root.Val + left + right
}
