import (
    "strconv"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    res := 0
    
    helper(root, "", &res)
    
    return res
}

func helper(root *TreeNode, numStr string, res *int) {
    if root == nil {
        return
    }
    
    numStr += strconv.Itoa(root.Val)
    
    if root.Left == nil && root.Right == nil {
        val, _ := strconv.Atoi(numStr)
        *res += val
    }
    
    if root.Left != nil {
        helper(root.Left, numStr, res)
    }
    if root.Right != nil {
        helper(root.Right, numStr, res)
    }
}
