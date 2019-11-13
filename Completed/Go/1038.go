/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func bstToGst(root *TreeNode) *TreeNode {
    m := 0
    return helper(root, &m)
}

func helper (root *TreeNode , m *int) *TreeNode {
    if root == nil { return root }
    helper(root.Right, m)
    
    root.Val += *m
    *m = root.Val
    
    helper(root.Left, m)
    return root
}