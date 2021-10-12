/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func equalToDescendants(root *TreeNode) int {
    res := 0
    
    helper(root, &res)
    
    return res
}

func helper(root *TreeNode, res *int) int {
    if root == nil {
        return 0
    }
    
    l := helper(root.Left, res)
    r := helper(root.Right, res)
    
    if l+r == root.Val {
        *res++
    }
    
    return root.Val+l+r
}
