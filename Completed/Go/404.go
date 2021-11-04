/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    res := 0
    
    helper(root, &res, false)
    
    return res
}

func helper(root *TreeNode, res *int, isLeft bool) {
    if root == nil {
        return
    }
    
    if root.Left == nil && root.Right == nil && isLeft {
        *res += root.Val
    }
    
    helper(root.Left, res, true)
    helper(root.Right, res, false)
}
