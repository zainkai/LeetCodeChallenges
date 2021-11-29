/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    return helper(root, root)
}

func helper(r1, r2 *TreeNode) bool {
    if r1 == nil && r2 == nil {
        return true
    } else if r1 == nil || r2 == nil {
        return false
    } else if r1.Val != r2.Val {
        return false
    }
    
    return helper(r1.Left, r2.Right) && helper(r1.Right, r2.Left)
}
